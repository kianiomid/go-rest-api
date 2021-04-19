package repository

import "api/models"

type UserRepository interface {
	//FindByEmail(email string) (*models.User, error)
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(id uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
}
