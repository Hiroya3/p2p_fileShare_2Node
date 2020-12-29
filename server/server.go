package server

import (
	"log"
	"net"
)

func StartServer() {
	//listenの開始
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("webサーバーを開始します")
}
