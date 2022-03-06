package tvshow_service

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/generated/models"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(tvShow models.TvShow) error {
	err := database.DB.Create(tvShow).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*models.TvShow, error) {
	var tvShows []*models.TvShow
	if result := database.DB.Find(&tvShows); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	fmt.Println(tvShows)
	return tvShows, nil
}
