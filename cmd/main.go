package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"p2p_fileShare_2Node/client"
	"p2p_fileShare_2Node/server"
)

type address struct {
	OwnDNS     string `json:"ownDNS"`
	OwnPort    string `json:"ownPort"`
	TargetDNS  string `json:"targetDNS"`
	TargetPort string `json:"targetPort"`
}

var Address address

//ノードの情報の読み取り
func init() {
	// nodeInfo/address.jsonから自ノードと相手ノードの情報を取得し、struct/address.goに格納
	jsonFile, err := ioutil.ReadFile("../nodeInfo/address.json")
	if err != nil {
		log.Fatalln("address.jsonの読み取りに失敗しました。エラー：%s", err)
	}

	err = json.Unmarshal(jsonFile, &Address)

	if err != nil {
		log.Fatalln("address.jsonの解析でエラーが発生しました。エラー：%s", err)
	}
}

func main() {
	/*
	* 処理の流れ
	* 1. 起動後listenerの起動をし他ノードからの検索を待つ
	* 2. 自ノードのコマンドラインにファイルの文言を表示
	 */

	//サーバーの開始
	go server.StartServer(Address.OwnDNS, Address.OwnPort)

	client.ViewCmd()
	fmt.Println("connに書き込みを行います")
	client.Query(Address.TargetDNS, Address.TargetPort)
	for {
	}
}
