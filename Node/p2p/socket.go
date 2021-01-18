package p2p

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
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
	buff := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(100 * time.Second))
	//forのなかでreaderを作ると2回目以降のループでreaderが再度初期化され何も残らなくなる
	//https://stackoverrun.com/ja/q/12701223
	reader := bufio.NewReader(conn)

	_, err := reader.Read(buff)
	if err != nil {
		if err == io.EOF {
			log.Fatalf("コネクションが閉じられました。\n エラー：%s", err)
		} else {
			log.Fatalf("検索ワードのコネクション読み込みでエラーが発生しました。\nエラー:%s", err)
		}
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
	sendSearchWords(connection, searchingWords)
}

//connetionに検索ワードを書き込む
//検索の場合は001:searchWords:[キーワード]:チェックサム（sha256）
func sendSearchWords(connection net.Conn, searchingWords []string) {
	if len(searchingWords) == 0 {
		log.Fatalln("キーワードを指定してください")
	}

	var messageBui strings.Builder

	messageBui.WriteString("001:searchWords:")
	messageBui.WriteString(strings.Join(searchingWords, ",") + ":")

	requestBodyStr := messageBui.String()

	//requestBodyStrのハッシュ値の計算
	sum := sha256.Sum256([]byte(requestBodyStr))
	requestStr := requestBodyStr + hex.EncodeToString(sum[:])

	_, err := connection.Write([]byte(requestStr))
	if err != nil {
		log.Fatalf("connectionへの書き込みに失敗しました。\nerr:%s", err)
	}
}
