package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func StartServer(address, port string) {
	//listenの開始
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("webサーバーを開始します")

	//Ctrl + Cで終了した時
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

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
			fmt.Println(string(messageBuff))
			fmt.Println("byte列の長さは" + string(messageLen) + "です")
		}()

		select {
		case <-interrupt:
			break
		}
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
