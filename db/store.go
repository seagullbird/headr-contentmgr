package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Store interface {
	InsertPost(post *Post) (id uint, err error)
	DeletePost(post *Post) error
	GetPost(id uint) (*Post, error)
	PatchPost(post *Post) error
	GetAllPosts(user_id string) ([]Post, error)
}

type databaseStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) Store {
	db.AutoMigrate(&Post{})
	return &databaseStore{
		db: db,
	}
}

func (s *databaseStore) InsertPost(post *Post) (id uint, err error) {
	err = s.db.Create(post).Error
	return post.Model.ID, err
}

func (s *databaseStore) DeletePost(post *Post) error {
	return s.db.Delete(&post).Error
}

func (s *databaseStore) GetPost(id uint) (*Post, error) {
	var post Post
	err := s.db.First(&post, id).Error
	return &post, err
}

func (s *databaseStore) PatchPost(post *Post) error {
	return nil
}

func (s *databaseStore) GetAllPosts(user_id string) ([]Post, error) {
	var posts []Post
	err := s.db.Where("user_id = ?", user_id).Find(&posts).Error
	return posts, err
}
