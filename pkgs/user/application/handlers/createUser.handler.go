package user_handlers

import (
	"encoding/json"
	"net/http"

	user_value_objects "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/valueObjects"
	commands_user "github.com/alianmc1988/go-restfull-gin/pkgs/user/repository/commands"
	"github.com/gin-gonic/gin"
)



func CreateUserHandler(c *gin.Context) {
	var userDTO user_value_objects.CreateUserDTO
	json.NewDecoder(c.Request.Body).Decode(&userDTO)
	data:= commands_user.CreateUser(&userDTO)
	userResponse:= user_value_objects.NewBaseUserResponseDTO(
		data.ID,
		data.FirstName,
		data.LastName,
		data.Email,
	)
	c.JSON(http.StatusCreated, userResponse)

}