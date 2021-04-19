package crud

import (
	"api/models"
	"api/utils/channels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
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
	go func(ch chan<- bool) {
		defer close(ch)

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

func (r *usersRepositoryCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)

		err := r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *usersRepositoryCRUD) FindById(uid uint32) (models.User, error) {
	var err error
	user := models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)
		err := r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err
}

func (r *usersRepositoryCRUD) Update(uid uint32, user models.User) (int64, error) {
	var rs *gorm.DB

	//user := models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumn(
			map[string]interface{}{
				"nickname": user.Nickname,
				"email": user.Email,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}

func (r *usersRepositoryCRUD) Delete(uid uint32) (int64, error) {
	var rs *gorm.DB

	//user := models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(models.User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}