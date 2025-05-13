package bfs

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"math"
	"sync"
	"time"
)

// func GetRecipeBFS(graph *types.RecipeGraph, target string, combos *[]types.Combo, nRecipe int, progressChan chan types.Combo, isLive bool) {
// 	temp := types.IngredientPair{target}
// 	queue := []types.IngredientPair{temp}
// 	parent := -1;

// 	var mu sync.Mutex
	
// 	var wg sync.WaitGroup
// 	for len(queue) > 0 {
// 		currentNode := queue[0]
// 		queue = queue[1:]
		
// 		for _, element := range currentNode{
			

// 			if graph.IsLeaf(element) || element == "" {
// 				continue
// 			}
			
// 			recipes := graph.Graph[element]
			
	
// 			listRecipe := graph.FilterTier(element)
	
// 			for i := 0; i < int(math.Min(float64(nRecipe),float64(len(listRecipe)))); i++{
	
// 				wg.Add(1)
	
// 				go func(rec types.IngredientPair, output string) {
// 					defer wg.Done()
					
// 					mu.Lock()
// 					if (!(graph.Graph[listRecipe[i][0]].Tier >= recipes.Tier || graph.Graph[listRecipe[i][1]].Tier >= recipes.Tier)){
// 						newCombo := types.Combo{
// 							ID:     len(*combos),
// 							ParentId: parent,
// 							Inputs: []string{rec[0], rec[1]},
// 							Output: output,
// 						}
// 						*combos = append(*combos, newCombo)
// 						mu.Unlock()
// 						mu.Lock()
						
// 						queue = append(queue, rec)
// 						nRecipe--
// 						mu.Unlock()
// 						if(isLive){
// 							progressChan <- newCombo

// 							time.Sleep(1000 * time.Millisecond)
// 						}
	
// 					}else{
// 						mu.Unlock()
// 					}
	
// 				}(listRecipe[i], element)
// 			}
// 			mu.Lock()
// 			nRecipe++
// 			mu.Unlock()
// 		}
// 		parent++
// 	}
// 	wg.Wait()
// }

func GetRecipeBFS(graph *types.RecipeGraph, target string, combos *[]types.Combo, nRecipe int, progressChan chan types.Combo, isLive bool) {
	var mu sync.Mutex
	queue := []types.IngredientPair{{target}}
	parent := -2

	for len(queue) > 0 {
		currentLevel := queue
		queue = nil
		
		var wg sync.WaitGroup
		for _, currentNode := range currentLevel {
			mu.Lock()
			parent++
			localParent := parent
			mu.Unlock()
			for _, element := range currentNode {
				if graph.IsLeaf(element) || element == "" {
					continue
				}

				recipes := graph.Graph[element]
				listRecipe := graph.FilterTier(element)
				mu.Lock()
				tempNRecipe := nRecipe
				mu.Unlock()
				if (tempNRecipe == 0){
					tempNRecipe = 1
				}
				limit := int(math.Min(float64(tempNRecipe), float64(len(listRecipe))))

				for i := 0; i < limit; i++ {
					rec := listRecipe[i]
					wg.Add(1)

					go func(rec types.IngredientPair, output string) {
						defer wg.Done()

						if graph.Graph[rec[0]].Tier >= recipes.Tier || graph.Graph[rec[1]].Tier >= recipes.Tier {
							return
						}

						newCombo := types.Combo{
							ID:       -1,
							ParentId: localParent,
							Inputs:   []string{rec[0], rec[1]},
							Output:   output,
						}
						
						mu.Lock()
						newCombo.ID = len(*combos)
						*combos = append(*combos, newCombo)
						queue = append(queue, rec)
						nRecipe--
						mu.Unlock()
						
						if isLive {
							progressChan <- newCombo
							time.Sleep(1 * time.Second)
						}
						}(rec, element)
					}
				}
				mu.Lock()
				nRecipe++
				mu.Unlock()
			}
		wg.Wait()
	}
}