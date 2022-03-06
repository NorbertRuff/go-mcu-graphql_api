package user_service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/norbertruff/go-graphql/graphql/generated/models"
	_ "github.com/norbertruff/go-graphql/graphql/generated/models"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"github.com/norbertruff/go-graphql/internal/pkg/utils"
	"log"
)

// Create user in db
func Create(user models.NewUser) (err error) {
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Println(err)
	}

	newUser := new(models.User)
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
func Authenticate(user models.User) bool {
	var row models.User
	result := database.DB.Where(&models.User{Username: user.Username}).First(&row)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return utils.CheckPasswordHash(user.Password, row.Password)
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	//statement, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//row := statement.QueryRow(username)
	//
	//var Id int
	//err = row.Scan(&Id)
	//if err != nil {
	//	if err != sql.ErrNoRows {
	//		log.Print(err)
	//	}
	//	return 0, err
	//}
	//
	//return Id, nil
	return 0, nil
}

// GetUsernameById check if a user exists in database and return the user object.
func GetUsernameById(userId string) (models.User, error) {
	//statement, err := database.Db.Prepare("select Username from Users WHERE ID = ?")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//row := statement.QueryRow(userId)
	//
	//var username string
	//err = row.Scan(&username)
	//if err != nil {
	//	if err != sql.ErrNoRows {
	//		log.Print(err)
	//	}
	//	return User{}, err
	//}
	//
	//return User{ID: userId, Username: username}, nil
	return models.User{}, nil
}
