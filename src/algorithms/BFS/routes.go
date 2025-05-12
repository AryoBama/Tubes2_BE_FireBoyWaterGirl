package bfs

import(
		
		// "fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"encoding/json"
	"unicode"
	"fmt"
	"time"
)

type Handler struct{
	Graph types.RecipeGraph
}

func NewHandler(graph types.RecipeGraph) *Handler{
	return &Handler{
		Graph: graph,
	}
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
		
		start := time.Now()
		GetRecipeBFS(&h.Graph, target, &combos, nRecipe)
		duration := time.Since(start)
		fmt.Printf("Waktunya bfs: %v", duration.Nanoseconds())

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"combos": combos,
			"duration": duration,
			"nNode" : len(combos),
		})
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