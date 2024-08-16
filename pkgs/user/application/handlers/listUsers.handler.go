package user_handlers

import (
	user_value_objects "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/valueObjects"
	users_queries "github.com/alianmc1988/go-restfull-gin/pkgs/user/repository/queries"
	"github.com/gin-gonic/gin"
)

func ListUsersHandler(c *gin.Context) {
	users:= users_queries.ListUsers()
	var usersResponse [] user_value_objects.BaseUserResponseDTO

	for _, user:= range users {
		nUser:= user_value_objects.NewBaseUserResponseDTO(
			user.ID,
			user.FirstName,
			user.LastName,
			user.Email,
		)
		usersResponse = append(usersResponse, *nUser)
	}

	c.JSON(200, usersResponse)
}