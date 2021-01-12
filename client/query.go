package client

import (
	"log"
	"net"
)

//検索の実行
func Query() {
	connection, err := net.Dial("tcp", "localhost:10000")

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

	for i := 0; i < len(searchingWords); i++ {
		_, err := connection.Write([]byte(searchingWords[i] + "\n"))
		if err != nil {
			log.Fatalf("Connectionへの書き込みでエラーが発生しました。\nerror:%s", err)
		}
	}
}
