package cmd

import (
	user_routes "github.com/alianmc1988/go-restfull-gin/pkgs/user/application/routes"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(s *Server) {
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	user_routes.InitializeUsersRoutes(s.Router)
}