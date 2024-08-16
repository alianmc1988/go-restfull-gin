package commands_user

import (
	"github.com/alianmc1988/go-restfull-gin/database"
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
	"github.com/gofrs/uuid"
)

func DeleteUserCommand(id string) string {

	var user user_models.User
	res:= database.DB.Where("id = ?", uuid.FromStringOrNil(id)).First(&user)
	if res.Error != nil {
		panic(res.Error)
	}

	database.DB.Delete(&user)

	return "User deleted successfully"

}