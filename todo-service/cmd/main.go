package main

import (
	"log"
	"todoservice/internal/cache"
	"todoservice/internal/database"
	"todoservice/internal/handler"
	"todoservice/internal/repository"
	"todoservice/internal/router"
	"todoservice/internal/service"
)

func main() {

	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	rdb, err := cache.ConnectToRedis()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	repo := repository.NewTodoRepo(db)
	if repo == nil {
		log.Fatal("failed to create repository")
	}

	s := service.NewTodoService(repo, rdb)
	if s == nil {
		log.Fatal("failed to create service")
	}

	h := handler.NewTodoHandler(s)

	r := router.NewRouter(h)

	const addr = "0.0.0.0:8082"

	log.Printf("✅ Auth service running on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
