package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
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
			log.Println("listenerのacceptでエラーが発生しました。err:%s", err)
		}
		go func() {
			defer conn.Close()
			// リクエストを読み込む
			messageBuff, messageLen := readRequestMessage(conn)
			fmt.Println(messageBuff)
			fmt.Println(messageLen)
		}()
	}
}

//リクエストを読み込む
func readRequestMessage(conn net.Conn) ([]byte, int) {
	messageBuff := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	messageLen, err := conn.Read(messageBuff)
	if err != nil {
		if err == io.EOF {
			//クライアント側から切断された時
			return []byte{}, 0
		}
		log.Fatalf("リクエストの読み込みに失敗しました。err:%s", err)
	}
	return messageBuff, messageLen
}
