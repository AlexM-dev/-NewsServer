package postgres

import (
	"GoNews/pkg/storage"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	password := "password"
	var dbURL string = fmt.Sprintf("postgres://postgres:" + password + "@65.108.250.159:5432/GoNews")
	_, err := New(dbURL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDB_News(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	posts := []storage.Post{
		{
			Title: "Lalala",
			Link:  strconv.Itoa(rand.Intn(353462347327)),
		},
	}
	password := "password"
	var dbURL string = fmt.Sprintf("postgres://postgres:" + password + "@65.108.250.159:5432/GoNews")
	db, err := New(dbURL)
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range posts {
		err = db.AddPost(i)
	}
	if err != nil {
		t.Fatal(err)
	}
	news, err := db.Posts(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", news)
}
