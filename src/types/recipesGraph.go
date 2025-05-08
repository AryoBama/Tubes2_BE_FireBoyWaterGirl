package types


type RecipeGraph struct {
    Recipes map[string][]IngredientPair
}

func NewRecipeGraph() *RecipeGraph {
    return &RecipeGraph{Recipes: make(map[string][]IngredientPair)}
}

func (g *RecipeGraph) AddRecipe(result string, ing1, ing2 string) {
    g.Recipes[result] = append(g.Recipes[result], IngredientPair{ing1, ing2})
}

func (g *RecipeGraph) ShowRecipes() string{
	result := ""

	for key,value := range g.Recipes{
		for _, ingredient := range value{
			
			result += key + " = " + ingredient[0] + " + " + ingredient[1] + "\n"
		}
	}

	return result
}

func (g *RecipeGraph) IsLeaf(element string) bool{

	return (element == "Air" || element == "Fire" || element == "Water" || element == "Earth")

}