package dfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	
)

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{
	}
}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func (h *Handler) HandleGetRecipe(router *mux.Router) {

	recipes := scrapper.ScrapRecipe()
	
	
	router.HandleFunc("/api/dfs", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(recipes.ShowRecipes()))
	}).Methods(http.MethodGet)


	router.HandleFunc("/api/dfs/{target}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		var combos []types.Combo
		temp := types.IngredientPair{target}
		nStr := r.URL.Query().Get("n")
		nRecipe:=1

 		updates := make(chan types.Combo)

		if nStr != "" {
			if val, err := strconv.Atoi(nStr); err == nil && val >= 0 {
				nRecipe = val - 1
			}
		}


		start := time.Now()
		GetRecipeDFS(recipes, temp, &combos, &nRecipe, updates, false)
		duration := time.Since(start)
	

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"combos": combos,
			"duration": duration.Microseconds(),
			"nNode" : len(combos),
		})

		// SaveRecipeTreeToFile(*tree,"test.json")
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/ws/dfs/{target}", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		var combos []types.Combo
		temp := types.IngredientPair{target}
		nStr := r.URL.Query().Get("n")
		nRecipe:=1

		updates := make(chan types.Combo)

		if nStr != "" {
			if val, err := strconv.Atoi(nStr); err == nil && val >= 0 {
				nRecipe = val - 1
			}
		}

		start := time.Now()
		go func() {
			GetRecipeDFS(recipes, temp, &combos, &nRecipe, updates, true)
		}()
		duration := time.Since(start)
		log.Println(duration)

		for combo := range updates {
			err := conn.WriteJSON(combo)
			if err != nil {
				log.Println("WebSocket write error:", err)
				break
			}
		}
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