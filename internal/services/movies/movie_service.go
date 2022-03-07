package movie_service

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/models"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(movie models.Movie) error {
	err := database.DB.Create(movie).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*models.Movie, error) {

	var movies []*models.Movie
	if result := database.DB.Find(&movies); result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(movies)
	return movies, nil
}
