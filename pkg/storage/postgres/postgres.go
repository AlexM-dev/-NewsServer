package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// New Конструктор объекта хранилища.
func New(c string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), c)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			author_id,
			title,
			content,
			created_at
		FROM  posts
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorID,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)
	}
	return posts, rows.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	_, err := s.db.Query(context.Background(), `
		INSERT INTO posts (title, content, author_id)
		VALUES ($1, $2, $3);
		`,
		p.Title,
		p.Content,
		p.AuthorID,
	)
	return err
}
func (s *Store) UpdatePost(p storage.Post) error {
	_, err := s.db.Query(context.Background(), `
		UPDATE posts
		SET title = $1, content = $2
		WHERE id = $3;
		`,
		p.Title,
		p.Content,
		p.ID,
	)
	return err
}
func (s *Store) DeletePost(p storage.Post) error {
	_, err := s.db.Query(context.Background(), `
		DELETE FROM posts WHERE id = $1;
		`,
		p.ID,
	)
	return err
}
