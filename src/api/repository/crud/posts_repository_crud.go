package crud

import (
	"api/models"
	"api/utils/channels"
	"github.com/jinzhu/gorm"
)

type postsRepositoryCRUD struct {
	db *gorm.DB
}

func NewPostSRepositoryCRUD(db *gorm.DB) *postsRepositoryCRUD {
	return &postsRepositoryCRUD{db}
}

func (r *postsRepositoryCRUD) FindAll() ([]models.Post, error) {
	var err error
	posts := []models.Post{}

	done := make(chan bool)

	go func(ch chan <- bool) {
		defer close(ch)

		err := r.db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}

		if len(posts) > 0 {
			for i, _ := range posts{
				err = r.db.Debug().Model(&models.Post{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		
		ch <- true

	}(done)

	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

func (r *postsRepositoryCRUD) Save(post models.Post) (models.Post, error)  {
	var err error

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)

		err := r.db.Debug().Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}
