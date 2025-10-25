package main

import (
	"authservice/internal/database"
	"authservice/internal/handler"
	"authservice/internal/repository"
	"authservice/internal/router"
	"authservice/internal/service"
	"log"
)

func main() {

	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewAuthRepo(db)
	if repo == nil {
		log.Fatal("failed to create repository")
	}

	s := service.NewAuthService(repo)
	if s == nil {
		log.Fatal("failed to create service")
	}

	h := handler.NewAuthHandler(s)

	r := router.NewRouter(h)

	log.Println("Starting server on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
