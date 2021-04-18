package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Commnet -
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// CommentService -
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newCommnet Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComment() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - retrieves comment by their ID from database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentBySlug - retrieves all comment bu slug (path - /article/name)
func (s *Service) getCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// UpdateComment - update a comment by ID with new comment info
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Update(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}
