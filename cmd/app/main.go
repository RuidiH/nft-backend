// backend/cmd/app/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/RuidiH/nft-backend/internal/db"
	"github.com/RuidiH/nft-backend/internal/handlers"
)

func main() {
	// Connect to the database
	db.ConnectDB()

	// Register HTTP handlers
	http.HandleFunc("/users", handlers.GetUsers)

	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
