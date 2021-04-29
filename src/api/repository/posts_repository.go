package repository

import "api/models"

type PostRepository interface {

	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
}
