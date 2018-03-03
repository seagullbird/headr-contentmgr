package db

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Store interface {
	InsertPost(post *Post) (id uint, err error)
	DeletePost(post *Post) error
	GetPost(id uint) (*Post, error)
	PatchPost(post *Post) error
}

type databaseStore struct {
	db *gorm.DB
}

func New(logger log.Logger) Store {
	db, err := gorm.Open("postgres", "host=postgresql-postgresql port=5432 user=postgres dbname=postgres password=qBDXNlz276 sslmode=disable")
	if err != nil {
		logger.Log("error_desc", "Failed to connected to PostgreSQL", "error", err)
	}
	db.AutoMigrate(&Post{})
	return &databaseStore{
		db: db,
	}
}

func (s *databaseStore) InsertPost(post *Post) (id uint, err error) {
	s.db.Create(post)
	return post.Model.ID, nil
}

func (s *databaseStore) DeletePost(post *Post) error {
	s.db.Delete(&post)
	return nil
}

func (s *databaseStore) GetPost(id uint) (*Post, error) {
	var post Post
	s.db.First(&post, id)
	return &post, nil
}

func (s *databaseStore) PatchPost(post *Post) error {
	return nil
}
