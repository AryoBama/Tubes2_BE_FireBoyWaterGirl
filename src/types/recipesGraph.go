package types

import(
	"fmt"
)

type Element struct{
	Name string `json:"name"`
	Recipes []IngredientPair `json:"recipes"`
	Tier int	`json:"tier"`
}

type RevElement struct{
	Output string `json:"output"`
	Partner string `json:"partner"`
}

type RecipeGraph struct {
    Graph map[string]Element
}

type ReverseGraph struct {
	Graph map[string][]RevElement
}


func NewRecipeGraph() *RecipeGraph {
    return &RecipeGraph{Graph: make(map[string]Element)}
}

func NewReverseGraph(graph *RecipeGraph) *ReverseGraph{
	result := ReverseGraph{Graph: make(map[string][]RevElement)}

	for _, element := range graph.Graph{
		for _, pair := range element.Recipes{
			result.Graph[pair[0]] = append(result.Graph[pair[0]], RevElement{Output: element.Name, Partner: pair[1]})
			result.Graph[pair[1]] = append(result.Graph[pair[1]], RevElement{Output: element.Name, Partner: pair[0]})
		}
	}
	return &result
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