package link_service

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/models"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
)

func Save(link *models.Link) error {
	err := database.DB.Create(link).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]*models.Link, error) {

	var links []*models.Link
	if result := database.DB.Find(&links); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	fmt.Println(links)
	return links, nil
}
