package main

import (
	bfs "Tubes2_BE_FireBoyWaterGirl/src/algorithms/BFS"
	dfs "Tubes2_BE_FireBoyWaterGirl/src/algorithms/DFS"
	live "Tubes2_BE_FireBoyWaterGirl/src/algorithms/Live"
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()


	data := scrapper.ScrapRecipe()

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	bfsHandler := bfs.NewHandler(*data)
	dfsHandler := dfs.NewHandler()
	dfsLiveHandler := live.NewHandler()

	bfsHandler.HandleGetRecipe(router)
	dfsHandler.HandleGetRecipe(router)
	dfsLiveHandler.HandleGetRecipe(router)


	router.HandleFunc("/api/elements", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		elements, err := scrapper.ScrapElements()
		if err != nil {
			log.Printf("Error scraping elements: %v", err)
			log.Printf("Falling back to base elements")
		}

		if len(elements) < 10 {
			log.Printf("Warning: Only scraped %d elements, might be incomplete", len(elements))
		}

		if err := json.NewEncoder(w).Encode(elements); err != nil {
			log.Printf("Error encoding elements to JSON: %v", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	fmt.Println("Server started at http://localhost:8080")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🔥 FireBoyWaterGirl API is running. Use /api, /api/bfs, or /api/dfs"))
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
