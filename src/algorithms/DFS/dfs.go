package dfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"

)



func GetRecipeDFS(graph *types.RecipeGraph, visited map[string]bool,current string, tree *types.RecipeNode) {
	
	if (visited[current] || graph.IsLeaf(current)){
		return
	}
	
	visited[current] = true

	pairs:= graph.Recipes[current]


	for _, pair := range pairs {
		ingredient1 := pair[0]
		ingredient2 := pair[1]

		
		if ingredient1 != "" {
			childNode1 := types.NewRecipeNode(ingredient1, graph.Recipes[ingredient1], []types.RecipeNode{})
			
			if !visited[ingredient1] {
				GetRecipeDFS(graph, visited, ingredient1, childNode1)
			}
			
			tree.AddChild(childNode1)
		}
		
		if ingredient2 != "" {
			childNode2 := types.NewRecipeNode(ingredient2, graph.Recipes[ingredient2], []types.RecipeNode{})
			
			if !visited[ingredient2] {
				GetRecipeDFS(graph, visited, ingredient2, childNode2)
			}
			
			tree.AddChild(childNode2)
		}
	}

}

func GetRecipeTree(graph *types.RecipeGraph, resultElement string) *types.RecipeTree {

	rootNode := types.NewRecipeNode(resultElement, graph.Recipes[resultElement], []types.RecipeNode{})
	
	visited := make(map[string]bool)
	
	GetRecipeDFS(graph, visited, resultElement, rootNode)
	
	return &types.RecipeTree{
		Root:   resultElement,
		Recipe: *rootNode,
	}
}