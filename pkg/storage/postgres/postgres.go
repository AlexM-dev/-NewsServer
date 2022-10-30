package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4"
)

// Store Хранилище данных.
type Store struct {
	db *pgx.Conn
}

// New Конструктор объекта хранилища.
func New(c string) (*Store, error) {
	db, err := pgx.Connect(context.Background(), c)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

func (s *Store) Posts(n int) ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			title,
			content,
			created_at,
			link
		FROM  posts
		WHERE id <= $1
		ORDER BY id;
	`,
		n,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PubTime,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)
	}
	return posts, rows.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		INSERT INTO posts (title, content, created_at, link)
		VALUES ($1, $2, $3, $4);
		`,
		p.Title,
		p.Content,
		p.PubTime,
		p.Link,
	)
	defer rows.Close()
	return err
}

func (s *Store) UpdatePost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		UPDATE posts
		SET title = $1, content = $2
		WHERE id = $3;
		`,
		p.Title,
		p.Content,
		p.ID,
	)
	defer rows.Close()
	return err
}
func (s *Store) DeletePost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		DELETE FROM posts WHERE id = $1;
		`,
		p.ID,
	)
	defer rows.Close()
	return err
}
