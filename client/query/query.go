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
	sendQuery(connection)
}

//connetionにqueryを書き込む
func sendQuery(connection net.Conn) {

}
