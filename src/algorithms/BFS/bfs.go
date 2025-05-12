package bfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"math"
	"sync"
)

func GetRecipeBFS(graph *types.RecipeGraph, target string, combos *[]types.Combo, nRecipe int) {
	temp := types.IngredientPair{target}
	queue := []types.IngredientPair{temp}
	parent := -1;

	var mu sync.Mutex
	
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		
		for _, element := range currentNode{
			

			if graph.IsLeaf(element) || element == "" {
				continue
			}
			
			recipes := graph.Graph[element]
			
			var wg sync.WaitGroup
	
			listRecipe := graph.FilterTier(element)
	
			for i := 0; i < int(math.Min(float64(nRecipe),float64(len(listRecipe)))); i++{
	
				wg.Add(1)
	
				go func(rec types.IngredientPair, output string) {
					defer wg.Done()
					
					mu.Lock()
					if (!(graph.Graph[listRecipe[i][0]].Tier >= recipes.Tier || graph.Graph[listRecipe[i][1]].Tier >= recipes.Tier)){
						*combos = append(*combos, types.Combo{
							ID:     len(*combos),
							ParentId: parent,
							Inputs: []string{rec[0], rec[1]},
							Output: output,
						})
						mu.Unlock()
						


						mu.Lock()
						
						queue = append(queue, rec)
						nRecipe--
						mu.Unlock()		

	
					}else{
						mu.Unlock()
					}
	
				}(listRecipe[i], element)
			}
			wg.Wait()
			nRecipe++
		}
		parent++
	}
}