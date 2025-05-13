package dfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"sync"
	"time"
)



func GetRecipeDFS(graph *types.RecipeGraph, currentNode types.IngredientPair, combos *[]types.Combo, nRecipe *int, progressChan chan types.Combo, isLive bool) {
	
	var mutex sync.Mutex
	var wg sync.WaitGroup

	dfsWithConcurrency(graph, currentNode, combos, nRecipe, &mutex, &wg, progressChan, isLive)

	wg.Wait()
	close(progressChan)
}

func dfsWithConcurrency(graph *types.RecipeGraph, currentNode types.IngredientPair, combos *[]types.Combo,
	nRecipe *int, mu *sync.Mutex, wg *sync.WaitGroup, progressChan chan types.Combo, isLive bool) {

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
			defer wg.Done()
			mu.Lock()
			comboID := len(*combos)
			newCombo := types.Combo{
				ID:       comboID,
				ParentId: parent,
				Inputs:   []string{p[0], p[1]},
				Output:   element,
			}
			*combos = append(*combos, newCombo)
			mu.Unlock()
			if(isLive){
				progressChan <- newCombo

				time.Sleep(1000 * time.Millisecond)
			}
			dfsWithConcurrency(graph, p, combos, nRecipe, mu, wg, progressChan, isLive)
			}(pair)
		}
	}
}