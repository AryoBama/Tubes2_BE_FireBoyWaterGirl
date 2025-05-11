package types

import(
	"fmt"
)

type Element struct{
	Name string `json:"name"`
	Recipes []IngredientPair `json:"recipes"`
	Tier int	`json:"tier"`
}

type RecipeGraph struct {
    Graph map[string]Element
}

func NewRecipeGraph() *RecipeGraph {
    return &RecipeGraph{Graph: make(map[string]Element)}
}

func (g *RecipeGraph) AddElement(result string, tier int){
	g.Graph[result] = Element{Name: result, Recipes: []IngredientPair{}, Tier: tier}
}

func (g *RecipeGraph) AddRecipe(result string, ing1, ing2 string) {
	temp := g.Graph[result]
	temp.Recipes = append(temp.Recipes, IngredientPair{ing1, ing2})
    g.Graph[result] = temp
}

func (g *RecipeGraph) FilterTier(target string) []IngredientPair{
	var result []IngredientPair
	for _, element := range g.Graph[target].Recipes{

		if(g.Graph[target].Tier <= g.Graph[element[0]].Tier || g.Graph[target].Tier <= g.Graph[element[1]].Tier){
			continue
		}

		result = append(result, element)
	}
	return result
}

func (g *RecipeGraph) ShowRecipes() string{
	result := ""

	for key,value := range g.Graph{
		for _, ingredient := range value.Recipes{
			
			result += key + " = " + ingredient[0] + " + " + ingredient[1]
		}
		result += "  Tier: " + fmt.Sprintf("%d", value.Tier) + "\n"
	}


	return result
}

func (g *RecipeGraph) IsLeaf(element string) bool{

	return (element == "Air" || element == "Fire" || element == "Water" || element == "Earth")

}