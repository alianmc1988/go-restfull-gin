package commands_user

import (
	"fmt"

	"github.com/alianmc1988/go-restfull-gin/database"
	user_entities "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/entities"
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
	user_value_objects "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/valueObjects"
)

func CreateUser(u *user_value_objects.CreateUserDTO) user_entities.UserEntity {

	userInstance:=user_models.NewUser(u)
	data:= database.DB.Create(&userInstance)
	if data.Error != nil {
		panic(fmt.Sprintf("Error: %v", data.Error))
	}
	userEntity:= user_entities.NewUserEntity(userInstance.ID.UUID.String(), userInstance.FirstName, userInstance.LastName, userInstance.Email, userInstance.Password)
	return *userEntity
}