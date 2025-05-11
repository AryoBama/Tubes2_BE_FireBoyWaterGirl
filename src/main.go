// main.go
package main

import (
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/BFS"
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/DFS"
	"Tubes2_BE_FireBoyWaterGirl/src/internal/scrapper"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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
	
	
	bfsHandler:=bfs.NewHandler(*data)
	dfsHandler:=dfs.NewHandler()
	
	
	bfsHandler.HandleGetRecipe(router)
	dfsHandler.HandleGetRecipe(router)
    handler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Bisa disesuaikan misal "http://localhost:3000"
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
    }).Handler(router)
	
	// for _,value := range data.Graph{
	// 	fmt.Println(value.Name + " " + "  Tier: " + fmt.Sprintf("%d", value.Tier))
	// 	fmt.Println(value.Recipes)
	// }
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

	
}