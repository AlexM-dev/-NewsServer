package storage

import (
	"strconv"
	"time"
)

// Post - публикация.
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

func CreatePost(id, authorID int, title, content string) *Post {
	t, _ := strconv.Atoi(time.Now().Format("20060102150405"))
	return &Post{
		ID:        id,
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: int64(t),
	}
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // Получение всех публикаций
	AddPost(Post) error     // Создание новой публикации
	UpdatePost(Post) error  // Обновление публикации
	DeletePost(Post) error  // Удаление публикации по ID
}
