package user_handlers

import (
	"fmt"
	"net/http"

	user_value_objects "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/valueObjects"
	commands_user "github.com/alianmc1988/go-restfull-gin/pkgs/user/repository/commands"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(c *gin.Context) {
	var userToUpdate = make(map[string]interface{})
	id:= c.Param("id")
	c.BindJSON(&userToUpdate)
	fmt.Println(userToUpdate)
	patchedUser,err:=commands_user.PatchUser(&id, &userToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	patchedUserSerialized:=user_value_objects.NewBaseUserResponseDTO(
		patchedUser.ID,
		patchedUser.FirstName,
		patchedUser.LastName,
		patchedUser.Email,
	)
	c.JSON(http.StatusOK, patchedUserSerialized)

}