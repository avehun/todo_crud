package app

import (
	"context"
	"log"
	"os"

	"github.com/avehun/todo_crud/internal/repo"
	"github.com/avehun/todo_crud/internal/server"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	log.Print("Starting server")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConnString := os.Getenv("DB_CONN_STRING")

	db, err := pgx.Connect(context.Background(), dbConnString)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close(context.Background())

	repo := repo.NewRepo(db)
	err = server.NewHandler(repo)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
