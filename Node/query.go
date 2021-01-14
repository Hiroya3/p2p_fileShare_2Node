package client

import (
	"log"
	"net"
	"strings"
)

//検索の実行
func Query(address, port string) {
	connection, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		log.Fatalf("query error!!! error:%s", err)
	}

	defer connection.Close()
	makeQuery(connection, SearchingWords)
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
