package main

import (
	"fmt"
	"log"
	"net/http"

	configs "github.com/SamuelSchutz13/SocialDev/config"
	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/routes"
)

func main() {
	conn, err := configs.NewConnectionDB()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer conn.Close()

	queries := db.New(conn)
	r := routes.SetupRoutes(queries)

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
