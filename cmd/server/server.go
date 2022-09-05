package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
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

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db1 := memdb.New()

	password := "password"

	var dbURL string = fmt.Sprintf("postgres://postgres:" + password + "@65.108.250.159:5432/GoNews")
	db2, err := postgres.New(dbURL)
	catchErr(err)
	// Документная БД MongoDB.
	db3, err := mongo.New("mongodb://localhost:27017/")
	catchErr(err)
	_, _, _ = db1, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db3

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	//err = db3.AddPost(*storage.CreatePost(9, 6, "New Post3", "Wow! New Post12!"))
	//catchErr(err)

	http.ListenAndServe(":8080", srv.api.Router())
}

func catchErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
