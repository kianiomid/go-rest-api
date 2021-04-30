package crud

import (
	"api/models"
	"api/utils/channels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
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

func (r *postsRepositoryCRUD) FindById(pid uint64) (models.Post, error) {
	var err error
	post := models.Post{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)

		err := r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&post).Error
		if err != nil {
			ch <- false
			return
		}

		if post.ID != 0 {
			err = r.db.Debug().Model(&models.Post{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.Post{}, errors.New("post not found")
	}
	return models.Post{}, err
}

func (r *postsRepositoryCRUD) Update(pid uint64, post models.Post) (int64, error) {
	var rs *gorm.DB

	//user := models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)

		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&models.Post{}).UpdateColumn(
			map[string]interface{}{
				"title": post.Title,
				"content": post.Content,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("Post not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}

func (r *postsRepositoryCRUD) Delete(pid uint64) (int64, error) {
	var rs *gorm.DB

	//user := models.User{}

	done := make(chan bool)

	//todo: check after this func
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&models.Post{}).Delete(models.Post{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("Post not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}
