package main

import (
	"fmt"
	"p2p_fileShare_2Node/client"
	"p2p_fileShare_2Node/server"
)

func main() {
	/*
	* 処理の流れ
	* 1. 起動後listenerの起動をし他ノードからの検索を待つ
	* 2. 自ノードのコマンドラインにファイルの文言を表示
	 */

	//サーバーの開始
	server.StartServer()

	//自ノードのコマンドラインにファイル検索の文言を表示
	fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。")
	searchingWords := client.GetSearchingWords()
	fmt.Println(searchingWords)
}
