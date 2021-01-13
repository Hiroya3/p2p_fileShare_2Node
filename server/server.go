package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func StartServer(address, port string) {
	//listenの開始
	listener, err := net.Listen("tcp", address+":"+port)
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
			messageSlice := readRequestMessage(conn)
			fmt.Println(messageSlice)
		}()
	}
}

//リクエストを読み込む
func readRequestMessage(conn net.Conn) []string {
	messageSlice := []string{}

	conn.SetReadDeadline(time.Now().Add(100 * time.Second))
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//最終行以降を読み込んだ際
				break
			}
			log.Fatalf("リクエストの読み込みに失敗しました。err:%s", err)
		}

		message = strings.ReplaceAll(message, "\n", "")
		messageSlice = append(messageSlice, message)
	}
	return messageSlice
}
