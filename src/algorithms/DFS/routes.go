package dfs

import (

	// "fmt"
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"encoding/json"
	"net/http"
	"unicode"
	"github.com/gorilla/mux"
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"os"
	"fmt"
)

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{
	}
}
func (h *Handler) HandleGetRecipe(router *mux.Router) {

	recipes := scrapper.ScrapRecipe()
	


	router.HandleFunc("/dfs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(recipes.ShowRecipes()))
	}).Methods(http.MethodGet)


	router.HandleFunc("/dfs/{target}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		target := string(unicode.ToUpper(rune(vars["target"][0]))) + vars["target"][1:]
		fmt.Println((*recipes).Recipes[target])
		cnt := 0
		tree := GetRecipeTree(recipes, target)
		
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Nodenya: " + fmt.Sprintf("%d", cnt)))
		// err := json.NewEncoder(w).Encode(tree)
		SaveRecipeTreeToFile(*tree,"test.json")

		// if err != nil {
		// 	http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		// }
		// response := GetRecipeBFS(recipe)

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