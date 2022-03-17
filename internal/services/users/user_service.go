package user_service

import (
	"github.com/google/uuid"
	"github.com/norbertruff/go-graphql/graphql/model"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"github.com/norbertruff/go-graphql/internal/pkg/utils"
	"log"
)

// Create user in db
func Create(user model.NewUser) (err error) {
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Println(err)
	}

	newUser := new(model.User)
	newUser.UserID = uuid.New().String()
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.Password = hashedPassword
	newUser.Firstname = user.Firstname
	newUser.Lastname = user.LastName
	newUser.Credit = 80
	err = database.DB.Create(newUser).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Authenticate check if user is in db and has a matching password
func Authenticate(user model.User) bool {
	var row model.User
	result := database.DB.Where(&model.User{Username: user.Username}).First(&row)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return utils.CheckPasswordHash(user.Password, row.Password)
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (string, error) {
	var row model.User
	result := database.DB.Where(&model.User{Username: username}).First(&row)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return row.UserID, nil

}

// GetUsernameById check if a user exists in database and return the user object.
func GetUsernameById(userId string) (string, error) {
	var row model.User
	result := database.DB.Where(&model.User{UserID: userId}).First(&row)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return row.Username, nil
}

// GetAll gets all users.
func GetAll() ([]*model.User, error) {
	var users []*model.User
	if result := database.DB.Find(&users); result.Error != nil {
		log.Println(result.Error)
	}
	return users, nil
}
