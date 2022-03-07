package link_service

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/model"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(link *model.Link) error {
	err := database.DB.Create(link).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*model.Link, error) {

	var links []*model.Link
	if result := database.DB.Find(&links); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return links, nil
}
