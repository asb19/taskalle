// @title User Service API
// @version 1.0
// @description This is a simple User microservice with CRUD, pagination, and filtering.
// @host localhost:8081
// @BasePath /
// @schemes http
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/asb19/usersvc/docs"
	"github.com/asb19/usersvc/internal/handler"
	"github.com/asb19/usersvc/internal/repo"
	"github.com/asb19/usersvc/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	// if dbURL == "" {
	// 	dbURL = "postgres://myuser:mypassword@localhost:5432/userdb?sslmode=disable"
	// }

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to user database: %v", err)
	}
	defer dbpool.Close()
	fmt.Println("✅ Connected to user database")

	repo := repo.NewPostgresUserRepository(dbpool)
	service := service.NewUserService(repo)
	userHandler := handler.NewHandler(service)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")

	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	fmt.Println("🚀 User Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
