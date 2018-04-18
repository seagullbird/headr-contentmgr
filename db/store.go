package db

import (
	"github.com/jinzhu/gorm"
	// used for database connection
	"errors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Store deals with database operations with table Post.
type Store interface {
	InsertPost(post *Post) (id uint, err error)
	DeletePost(post *Post) error
	GetPost(id uint) (*Post, error)
	PatchPost(existing, post *Post) (*Post, error)
	GetAllPosts(userID string) ([]uint, error)
}

type databaseStore struct {
	db *gorm.DB
}

// New creates a databaseStore instance
func New(db *gorm.DB) Store {
	db.AutoMigrate(&Post{})
	return &databaseStore{
		db: db,
	}
}

func (s *databaseStore) InsertPost(post *Post) (id uint, err error) {
	var p []Post
	s.db.Where("title = ?", post.Title).Find(&p)
	if len(p) > 0 {
		return 0, errors.New("title already exists")
	}
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

func (s *databaseStore) PatchPost(existing, post *Post) (*Post, error) {
	if post.Title != "" {
		existing.Title = post.Title
	}
	if post.Tags != "" {
		existing.Tags = post.Tags
	}
	if post.Draft != existing.Draft {
		existing.Draft = post.Draft
	}
	if post.Date != "" {
		existing.Date = post.Date
	}
	if post.Summary != "" {
		existing.Summary = post.Summary
	}
	err := s.db.Save(existing).Error
	return existing, err
}

func (s *databaseStore) GetAllPosts(userID string) ([]uint, error) {
	var posts []Post
	err := s.db.Where("user_id = ?", userID).Find(&posts).Error
	var ids = make([]uint, len(posts))
	for i, v := range posts {
		ids[i] = v.Model.ID
	}
	return ids, err
}
