package live

import (
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) HandleGetRecipe(router *mux.Router) {
	recipes := scrapper.ScrapRecipe()

	router.HandleFunc("/api/dfs/live/{target}", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		temp := types.IngredientPair{target}

		nRecipe := 1
		if nStr := r.URL.Query().Get("n"); nStr != "" {
			if val, err := strconv.Atoi(nStr); err == nil && val > 0 {
				nRecipe = val
			}
		}

		updates := make(chan types.Combo)
		start := time.Now()

		go GetRecipeDFS(recipes, temp, &[]types.Combo{}, &nRecipe, updates)

		log.Println("🔁 Live DFS dimulai untuk:", target)

		for combo := range updates {
			log.Println("📤 Kirim combo ke frontend:", combo)
			err := conn.WriteJSON(combo)
			if err != nil {
				log.Println("WebSocket write error:", err)
				break
			}
		}

		log.Println("⏱️ Durasi pencarian:", time.Since(start))
	}).Methods(http.MethodGet)
}

func SaveRecipeTreeToFile(tree types.RecipeTree, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}
