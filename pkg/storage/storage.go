package storage

import (
	"strconv"
	"time"
)

var c int = 0

// Post Публикация, получаемая из RSS.
type Post struct {
	ID      int    // Номер записи
	Title   string // Заголовок публикации
	Content string // Содержание публикации
	PubTime int64  // Время публикации
	Link    string // Ссылка на источник
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts(n int) ([]Post, error) // Получение всех публикаций
	AddPost(Post) error          // Создание новой публикации
	UpdatePost(Post) error       // Обновление публикации
	DeletePost(Post) error       // Удаление публикации по ID
}

func CreatePost(title, content, link string) *Post {
	t, _ := strconv.Atoi(time.Now().Format("20060102150405"))
	temp := &Post{
		ID:      c,
		Title:   title,
		Content: content,
		Link:    link,
		PubTime: int64(t),
	}
	c++
	return temp
}
