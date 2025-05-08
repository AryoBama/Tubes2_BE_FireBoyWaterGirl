// main.go
package main

import (
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/BFS"
	"Tubes2_BE_FireBoyWaterGirl/src/algorithms/DFS"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	bfsHandler:=bfs.NewHandler()
	dfsHandler:=dfs.NewHandler()


	bfsHandler.HandleGetRecipe(router)
	dfsHandler.HandleGetRecipe(router)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	
}