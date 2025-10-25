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

	svc := service.NewAuthService(repo)
	if svc == nil {
		log.Fatal("failed to create service")
	}

	h := handler.NewAuthHandler(svc)

	r := router.NewRouter(h)

	const addr = "0.0.0.0:8081"

	log.Printf("✅ Auth service running on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
