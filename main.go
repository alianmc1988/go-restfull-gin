package main

import (
	"github.com/alianmc1988/go-restfull-gin/base"
	"github.com/alianmc1988/go-restfull-gin/cmd"
)


func main() {
	config:= base.Config


	s:= cmd.NewServer(config["HOST"], config["PORT"])
	s.Run()

}