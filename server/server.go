package server

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func StartServer() {
	//listenの開始
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("webサーバーを開始します")

	//1回でlistenerが閉じてしまわないようにfor文かつgoroutineで回す
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listenerのacceptでエラーが発生しました")
			log.Fatalln(err)
		}
		go func() {
			//レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
