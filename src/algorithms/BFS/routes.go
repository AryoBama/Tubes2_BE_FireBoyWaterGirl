package bfs

import(
		
		// "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"encoding/json"
	"unicode"
)

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{
	}
}
func (h *Handler) HandleGetRecipe(router *mux.Router) {

	recipes := scrapper.ScrapRecipe()

	router.HandleFunc("/bfs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(recipes.ShowRecipes()))
	}).Methods(http.MethodGet)


	router.HandleFunc("/bfs/{target}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		var combos []types.Combo

		GetRecipeBFS(recipes, target, &combos)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("bfs" + target + "\n"))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"combos": combos,
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