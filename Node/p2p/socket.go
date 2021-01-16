package p2p

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func Run(address, port string) {
	//listenの開始
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
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
	//forのなかでreaderを作ると2回目以降のループでreaderが再度初期化され何も残らなくなる
	//https://stackoverrun.com/ja/q/12701223
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

//相手ノードへのファイル検索の実行
func SearchFile(address, port string, searchingWords []string) {
	connection, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		log.Fatalf("query error!!! error:%s", err)
	}

	defer connection.Close()
	makeQuery(connection, searchingWords)
}

//connetionにqueryを書き込む
func makeQuery(connection net.Conn, searchingWords []string) {
	if len(searchingWords) == 0 {
		log.Fatalln("キーワードを指定してください")
	}

	//server側で改行で取得するため末尾にも改行コードを入れる
	_, err := connection.Write([]byte(strings.Join(searchingWords, "\n") + "\n"))
	if err != nil {
		log.Fatalf("Connectionへの書き込みでエラーが発生しました。\nerror:%s", err)
	}
}
