package repository

import "projects/go-rest-api/src/api/models"

type UserRepository interface {
	//FindAll() ([]models.User, error)
	//FindById(id int32) (models.User, error)
	//FindByEmail(email string) (*models.User, error)

	Save(models.User) (models.User, error)
	//Update(int32, models.User) (int64, error)
	//Delete(int32) (int64, error)
}
