package main

import (
	"fmt"
	"log"
	"net/http"

	"my-api/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
