package main

import (
	"GoNews/internal/app"
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres"
	"fmt"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	password := "password"
	var dbURL string = fmt.Sprintf("postgres://postgres:" + password + "@65.108.250.159:5432/GoNews")
	db, err := postgres.New(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	go app.Run(db)

	http.ListenAndServe(":80", srv.api.Router())
}
