package main

import (
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/BFS"
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/DFS"
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	// "Tubes2_BE_FireBoyWaterGirl/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Membuat router baru
	router := mux.NewRouter()

	data := scrapper.ScrapRecipe()
	// reverseData := types.NewReverseGraph(data)

	// for target, combination := range data.Graph{
	// 	for _, recipe := range combination.Recipes{
	// 		if (data.Graph[recipe[0]].Tier + 1 < data.Graph[target].Tier && data.Graph[recipe[1]].Tier + 1 < data.Graph[target].Tier){
	// 			fmt.Printf("Target: %s\n", target)
	// 			fmt.Printf("rec 1: %s,  rec 2: %s\n",recipe[0],recipe[1])
	// 		}
	// 	}
	// }
	
	// Route untuk API utama
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	bfsHandler := bfs.NewHandler(*data)
	dfsHandler := dfs.NewHandler()

	bfsHandler.HandleGetRecipe(router)
	dfsHandler.HandleGetRecipe(router)

	router.HandleFunc("/api/elements", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Scrape elemen
		elements, err := scrapper.ScrapElements()
		if err != nil {
			log.Printf("Error scraping elements: %v", err)
			log.Printf("Falling back to base elements")
			// elements = scrapper.GetBaseElements()
		}

		if len(elements) < 10 {
			log.Printf("Warning: Only scraped %d elements, might be incomplete", len(elements))
		}

		// Kembalikan data elemen dalam format JSON
		if err := json.NewEncoder(w).Encode(elements); err != nil {
			log.Printf("Error encoding elements to JSON: %v", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Setup CORS handler
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Sesuaikan dengan URL frontend jika perlu
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	// Informasi server sudah berjalan
	fmt.Println("Server started at http://localhost:8080")

	// Handle route untuk status
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🔥 FireBoyWaterGirl API is running. Use /api, /api/bfs, or /api/dfs"))
	})

	// Jalankan server pada port 8080
	log.Fatal(http.ListenAndServe(":8080", handler))
}
