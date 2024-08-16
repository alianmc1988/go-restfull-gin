package main

import (
	"github.com/alianmc1988/go-restfull-gin/cmd"
)

const PORT = "8080"
const HOST = "localhost"

func main() {
	s:= cmd.NewServer(HOST, PORT)
	s.Run()

}