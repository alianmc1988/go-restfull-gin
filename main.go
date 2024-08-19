package main

import (
	"github.com/alianmc1988/go-restfull-gin/base"
	"github.com/alianmc1988/go-restfull-gin/cmd"
)


func main() {
	config:= base.GetConfig()


	s:= cmd.NewServer(config.Host, config.Port)
	s.Run()

}