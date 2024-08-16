package commands_user

import (
	"github.com/alianmc1988/go-restfull-gin/database"
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
)

func DeleteUserCommand(id string) string {

	var user user_models.User
	res:= database.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		panic(res.Error)
	}

	err:= database.DB.Delete(&user)
	if err.Error != nil {
		panic(err.Error)
	}

	return "User deleted successfully"

}