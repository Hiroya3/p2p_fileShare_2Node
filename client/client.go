package client

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var SearchingWords []string

//コマンドラインに案内画面を表示します。
func ViewCmd() {
	for {
		fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。\n３つ以上指定した場合ははじめ３つが採用されます。")
		SearchingWords = getSearchingWords(os.Stdin)

		if len(SearchingWords) > 3 {
			fmt.Println("4つ以上指定されたため、はじめの３つを取得します。")
			SearchingWords = SearchingWords[:3]
		}
		if len(SearchingWords) > 0 {
			break
		}
		if len(SearchingWords) == 0 {
			fmt.Println("キーワードを指定してください。")
		}
	}
}

//標準入力されたワードをスペース区切りで取得し、スライスに格納します
//unitテストを可能にするために引数にio.Readerを用いる
func getSearchingWords(stdin io.Reader) []string {
	var searchingWords = []string{}

	sc := bufio.NewScanner(stdin)
	sc.Scan()
	keywords := sc.Text()
	searchingWords = strings.Fields(keywords)

	return searchingWords
}
