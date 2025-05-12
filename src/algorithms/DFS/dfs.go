package dfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
)



func GetRecipeDFS(graph *types.RecipeGraph, currentNode types.IngredientPair, combos *[]types.Combo) {

	
	parent := len(*combos) - 1

	for _, element := range currentNode{

		pairs:= graph.FilterTier(element)
	
		for _, pair := range pairs {
			ingredient1 := pair[0]
			ingredient2 := pair[1]
	
			if ingredient1 != "" && ingredient2 != "" {
				*combos = append(*combos, types.Combo{
				ID: len(*combos),
				ParentId: parent,
				Inputs: []string{ingredient1, ingredient2},
				Output: element,
				})
				GetRecipeDFS(graph, pair, combos)		
			}
		}
	}

}