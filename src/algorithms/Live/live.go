package live

import (
	"Tubes2_BE_FireBoyWaterGirl/src/types"
	"sync"
	"time"
)

// Fungsi utama untuk DFS dengan live update via channel
func GetRecipeDFS(
	graph *types.RecipeGraph,
	currentNode types.IngredientPair,
	combos *[]types.Combo,
	nRecipe *int,
	progressChan chan types.Combo,
) {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	go func() {
		dfsWithConcurrency(graph, currentNode, combos, nRecipe, &mutex, &wg, progressChan)
		wg.Wait()
		close(progressChan) // channel hanya ditutup sekali di sini
	}()
}

func dfsWithConcurrency(
	graph *types.RecipeGraph,
	currentNode types.IngredientPair,
	combos *[]types.Combo,
	nRecipe *int,
	mu *sync.Mutex,
	wg *sync.WaitGroup,
	progressChan chan types.Combo,
) {
	parent := len(*combos) - 1

	for _, element := range currentNode {
		if graph.IsLeaf(element) || element == "" {
			continue
		}

		mu.Lock()
		*nRecipe++
		mu.Unlock()

		pairs := graph.FilterTier(element)

		for _, pair := range pairs {
			mu.Lock()
			if *nRecipe <= 0 {
				mu.Unlock()
				continue
			}
			*nRecipe--
			mu.Unlock()

			wg.Add(1)
			go func(p types.IngredientPair) {
				defer wg.Done()

				mu.Lock()
				newCombo := types.Combo{
					ID:       len(*combos),
					ParentId: parent,
					Inputs:   []string{p[0], p[1]},
					Output:   element,
				}
				*combos = append(*combos, newCombo)
				mu.Unlock()

				progressChan <- newCombo
				time.Sleep(1 * time.Second)

				dfsWithConcurrency(graph, p, combos, nRecipe, mu, wg, progressChan)
			}(pair)
		}
	}
}
