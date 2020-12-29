package main

//go.modでローカルをインポートするにはgomodのリプレイスを使用する
//https://qiita.com/hnishi/items/a9217249d7832ed2c035
import "local.packages/server"

func main() {
	/*
	* 処理の流れ
	* 1. 起動後listenerの起動をし他ノードからの検索を待つ
	* 2. 自ノードのコマンドラインにファイルの文言を表示
	 */

	//サーバーの開始
	server.StartServer()
}
