package tvshow_service

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/model"

	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(tvShow *model.TvShow) error {
	err := database.DB.Create(tvShow).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*model.TvShow, error) {
	var tvShows []*model.TvShow
	if result := database.DB.Find(&tvShows); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return tvShows, nil
}
