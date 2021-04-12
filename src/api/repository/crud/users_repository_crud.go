package crud

import (
	"github.com/jinzhu/gorm"
	"projects/go-rest-api/src/api/models"
	"projects/go-rest-api/src/api/utils/channels"
)

type usersRepositoryCRUD struct {
	db *gorm.DB
}

func NewUsersRepositoryCRUD(db *gorm.DB) *usersRepositoryCRUD {
	return &usersRepositoryCRUD{db}
}

func (r *usersRepositoryCRUD) Save(user models.User) (models.User, error) {
	var err error

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan <- bool) {
		err := r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}
