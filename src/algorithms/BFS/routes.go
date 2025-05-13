package bfs

import(
		
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"encoding/json"
	"unicode"
	"time"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
)

type Handler struct{
	Graph types.RecipeGraph
}

func NewHandler(graph types.RecipeGraph) *Handler{
	return &Handler{
		Graph: graph,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func (h *Handler) HandleGetRecipe(router *mux.Router) {

	router.HandleFunc("/api/bfs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(h.Graph.ShowRecipes()))
	}).Methods(http.MethodGet)


	router.HandleFunc("/api/bfs/{target}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		var combos []types.Combo

		
		nStr := r.URL.Query().Get("n")
		nRecipe:=1

		if nStr != "" {
			if val, err := strconv.Atoi(nStr); err == nil && val >= 0 {
				nRecipe = val
			}
		}

		updates := make(chan types.Combo)
		
		start := time.Now()
		GetRecipeBFS(&h.Graph, target, &combos, nRecipe, updates, false)
		close(updates)
		duration := time.Since(start)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"combos": combos,
			"duration": duration.Microseconds(),
			"nNode" : len(combos),
		})
	}).Methods(http.MethodGet)

		router.HandleFunc("/api/ws/bfs/{target}", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		var combos []types.Combo
		nStr := r.URL.Query().Get("n")
		nRecipe:=1

		updates := make(chan types.Combo)

		if nStr != "" {
			if val, err := strconv.Atoi(nStr); err == nil && val >= 0 {
				nRecipe = val
			}
		}

		start := time.Now()
		go func() {
			GetRecipeBFS(&h.Graph,target, &combos, nRecipe, updates, true)
			close(updates)
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

// func SaveRecipeTreeToFile(combo []types.Combo, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ")
// 	return encoder.Encode(tree)
// }