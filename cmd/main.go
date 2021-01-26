package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"p2p_fileShare_2Node/Node/p2p"
	"p2p_fileShare_2Node/Node/service"
)

type address struct {
	OwnDNS     string `json:"ownDNS"`
	OwnPort    string `json:"ownPort"`
	TargetDNS  string `json:"targetDNS"`
	TargetPort string `json:"targetPort"`
}

// Address is the DNS and Port info for ownNode and targetNode
var Address address

//ノードの情報の読み取り
func init() {
	// nodeInfo/address.jsonから自ノードと相手ノードの情報を取得し、struct/address.goに格納
	jsonFile, err := ioutil.ReadFile("../nodeInfo/address.json")
	if err != nil {
		log.Fatalf("address.jsonの読み取りに失敗しました。エラー：%s", err)
	}

	err = json.Unmarshal(jsonFile, &Address)

	if err != nil {
		log.Fatalf("address.jsonの解析でエラーが発生しました。エラー：%s", err)
	}
}

func main() {
	/*
	* 処理の流れ
	* 1. 起動後listenerの起動をし他ノードからの検索を待つ
	* 2. 自ノードのコマンドラインにファイルの文言を表示
	 */

	//サーバーの開始
	go p2p.Run(Address.OwnDNS, Address.OwnPort)

	searchingWords := service.GetSearchingWords()
	fmt.Println("connに書き込みを行います")
	searchedFiles, err := p2p.SearchFile(Address.TargetDNS, Address.TargetPort, searchingWords)
	if err != nil {
		fmt.Println("検索に失敗しました。時間をおいて再度実行してみて下さい。")
		log.Fatalln(err)

	}
	fmt.Println(service.SelectFileName(searchedFiles))
	for {
	}
}
