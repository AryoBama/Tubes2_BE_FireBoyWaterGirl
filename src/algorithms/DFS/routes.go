package dfs

import (

	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"encoding/json"
	"net/http"
	"unicode"
	"github.com/gorilla/mux"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"os"
	// "fmt"
)

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{
	}
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
		GetRecipeDFS(recipes, temp, &combos)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"combos": combos,
		})

		// SaveRecipeTreeToFile(*tree,"test.json")
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