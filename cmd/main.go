package main

import (
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
	go server.StartServer()

	client.ViewCmd()
}
