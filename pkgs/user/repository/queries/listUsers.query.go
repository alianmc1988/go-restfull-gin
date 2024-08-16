package users_queries

import (
	"github.com/alianmc1988/go-restfull-gin/database"
	user_entities "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/entities"
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
)

func ListUsers() ( [] user_entities.UserEntity) {
	var users [] user_models.User
	datas:=database.DB.Find(&users)
	if datas.Error != nil {
		panic(datas.Error)
	}
	var usersEntities [] user_entities.UserEntity

	for _, user:= range users {
		nUser:= user_entities.NewUserEntity(user.ID, user.FirstName, user.LastName, user.Email, user.Password)
		usersEntities = append(usersEntities, *nUser)
	}

	return usersEntities
	
}