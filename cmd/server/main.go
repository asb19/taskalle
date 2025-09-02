package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/asb19/tasksvc/internal/handler"
	"github.com/asb19/tasksvc/internal/repo"
	"github.com/asb19/tasksvc/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	}

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	} else {
		fmt.Println("connected to database")
	}
	defer dbpool.Close()
	repo := repo.NewPostgresTaskRepository(dbpool)
	service := service.NewTaskService(repo)
	taskHandler := handler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskHandler.GetAllTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	fmt.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
