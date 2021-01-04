package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var SearchingWords []string

//コマンドラインに案内画面を表示します。
func ViewCmd() {
	for {
		fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。\n３つ以上指定した場合ははじめ３つが採用されます。")
		SearchingWords = getSearchingWords()

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

//検索ワードを3つ取得します。4つ以上入力された場合ははじめの3つが取得されます。
func getSearchingWords() []string {
	var searchingWords = []string{}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	keywords := sc.Text()
	searchingWords = strings.Fields(keywords)

	return searchingWords
}
