package user_handlers

import (
	commands_user "github.com/alianmc1988/go-restfull-gin/pkgs/user/repository/commands"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	id:= c.Param("id")
	data:= commands_user.DeleteUserCommand(id)
	c.JSON(200, gin.H{
		"message": data,
	})
}