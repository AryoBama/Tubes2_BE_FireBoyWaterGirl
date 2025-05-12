package dfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"sync"
)



func GetRecipeDFS(graph *types.RecipeGraph, currentNode types.IngredientPair, combos *[]types.Combo, nRecipe *int) {
	
	// parent := len(*combos) - 1
	
	// for _, element := range currentNode{
	// 	if (graph.IsLeaf(element) || element == ""){
	// 		continue

	// 	}
	// 	*nRecipe+=1
		
		
	// 	pairs:= graph.FilterTier(element)
	// 	for _, pair := range pairs {
	// 		if(*nRecipe <= 0){
	// 			continue
	// 		}
	// 		fmt.Println(pair)
			
	// 		ingredient1 := pair[0]
	// 		ingredient2 := pair[1]
			
	// 		(*nRecipe)--
	// 		if ingredient1 != "" && ingredient2 != "" {
	// 			*combos = append(*combos, types.Combo{
	// 				ID: len(*combos),
	// 				ParentId: parent,
	// 				Inputs: []string{ingredient1, ingredient2},
	// 				Output: element,
	// 			})
	// 			GetRecipeDFS(graph, pair, combos, nRecipe)		
	// 		}
	// 	}
	// }
	var mutex sync.Mutex
	var wg sync.WaitGroup

	dfsWithConcurrency(graph, currentNode, combos, nRecipe, &mutex, &wg)

	wg.Wait()
}

func dfsWithConcurrency(graph *types.RecipeGraph, currentNode types.IngredientPair, combos *[]types.Combo, 
                        nRecipe *int, mu *sync.Mutex, wg *sync.WaitGroup) {

	parent := len(*combos) - 1

	for _, element := range currentNode {
		if graph.IsLeaf(element) || element == ""{
			continue
		}
		
		mu.Lock()
		*nRecipe += 1

		mu.Unlock()
		
		
		pairs := graph.FilterTier(element)
		
		for _, pair := range pairs {
			mu.Lock()
			if *nRecipe <= 0 {
				mu.Unlock()
				continue
			}
			
			(*nRecipe)--
			mu.Unlock()
			wg.Add(1)
			go func(p types.IngredientPair) {
			mu.Lock()
			comboID := len(*combos)
			*combos = append(*combos, types.Combo{
				ID:       comboID,
				ParentId: parent,
				Inputs:   []string{p[0], p[1]},
				Output:   element,
			})
			mu.Unlock()
			defer wg.Done()
			dfsWithConcurrency(graph, p, combos, nRecipe, mu, wg)
			}(pair)
		}
	}
}