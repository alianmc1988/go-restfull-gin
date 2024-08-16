package cmd

import (
	"github.com/alianmc1988/go-restfull-gin/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Host string
	Port string
	Router *gin.Engine
}

func NewServer(host string, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
		Router: gin.Default(),
	}
}

func (s *Server) Initialize(){
	DB:=database.DBConnect()
	database.MigrateModels(DB)

	InitializeRoutes(s)	
}

func (s *Server) Run() {
	s.Initialize()
	s.Router.Run(s.Host + ":" + s.Port)

}