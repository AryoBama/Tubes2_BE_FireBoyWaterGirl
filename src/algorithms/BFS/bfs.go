package bfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"sync"
)

func GetRecipeBFS(graph *types.RecipeGraph, target string, combos *[]types.Combo) {
	queue := []string{target}
	
	inQueue := make(map[string]bool)
	inQueue[target] = true
	
	processedCombos := make(map[string]bool)
	
	var mu sync.Mutex
	
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		
		mu.Lock()
		inQueue[currentNode] = false
		mu.Unlock()
		
		if graph.IsLeaf(currentNode) {
			continue
		}

		recipes := graph.Recipes[currentNode]
		
		var wg sync.WaitGroup
		for _, recipe := range recipes {
			wg.Add(1)
			
			go func(rec types.IngredientPair, output string) {
				defer wg.Done()
				
				comboKey := rec[0] + "," + rec[1] + "->" + output
				
				mu.Lock()
				if processedCombos[comboKey] {
					mu.Unlock()
					return
				}
				
				processedCombos[comboKey] = true
				
				*combos = append(*combos, types.Combo{
					ID:     len(*combos),
					Inputs: []string{rec[0], rec[1]},
					Output: output,
				})
				mu.Unlock()
				
				for i := 0; i < 2; i++ {
					ingredient := rec[i]
					
					if graph.IsLeaf(ingredient) {
						continue
					}
					
					mu.Lock()
					if !inQueue[ingredient] {
						queue = append(queue, ingredient)
						inQueue[ingredient] = true
					}
					mu.Unlock()
				}
			}(recipe, currentNode)
		}
		
		wg.Wait()
	}
}