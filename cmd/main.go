package main

import (
	"log"

	"github.com/danh996/video-chat/internal/server"
)


func main(){
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}

	
}