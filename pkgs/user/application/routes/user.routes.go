package user_routes

import (
	"fmt"
	"net/http"

	"github.com/alianmc1988/go-restfull-gin/base"
	user_handlers "github.com/alianmc1988/go-restfull-gin/pkgs/user/application/handlers"
	"github.com/gin-gonic/gin"
)

const BaseUserPath = "/user"


var userRoutes = []base.RouterType{
	{
		Path: BaseUserPath, 
		Method: http.MethodPost, 
		Handler: user_handlers.CreateUserHandler,
	},
	{
		Path: BaseUserPath,
		Method: http.MethodGet,
		Handler: user_handlers.ListUsersHandler,
	},
	{
		Path: fmt.Sprintf("%s/:id", BaseUserPath),
		Method: http.MethodDelete,
		Handler: user_handlers.DeleteUserHandler,
	},
}



func InitializeUsersRoutes(r *gin.Engine) *gin.Engine {
	for _, route := range userRoutes {
		r.Handle(route.Method, route.Path, route.Handler)
	}
	return r
}