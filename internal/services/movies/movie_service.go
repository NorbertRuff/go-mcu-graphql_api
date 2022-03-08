package movie_service

import (
	"github.com/norbertruff/go-graphql/graphql/model"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(movie *model.Movie) error {
	err := database.DB.Create(movie).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*model.Movie, error) {

	var movies []*model.Movie
	if result := database.DB.Find(&movies); result.Error != nil {
		log.Println(result.Error)
	}

	return movies, nil
}
