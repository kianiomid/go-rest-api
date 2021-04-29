package auto

import (
	"api/database"
	"api/models"
	"log"
)

func Load()  {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Post{} ,&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.Post{} ,&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	//add foreign keys
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		// add foreign key
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		/*err = db.Debug().Model(&posts[i]).Related(&posts[i].Author).Error
		if err != nil {
			log.Fatal(err)
		}*/

		//console.Pretty(posts[i])
	}
}
