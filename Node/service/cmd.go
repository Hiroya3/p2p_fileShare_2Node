package service

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//コマンドラインに検索を表示します。
func GetSearchingWords() []string {
	searchingWords := []string{}
	for {
		fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。\n３つ以上指定した場合ははじめ３つが採用されます。")
		//SearchingWords = getSearchingWordsBySpace(os.Stdin)
		searchingWords = []string{"aa", "bb", "cc"}

		if len(searchingWords) > 3 {
			fmt.Println("4つ以上指定されたため、はじめの３つを取得します。")
			searchingWords = searchingWords[:3]
		}
		if len(searchingWords) > 0 {
			break
		}
		if len(searchingWords) == 0 {
			fmt.Println("キーワードを指定してください。")
		}
	}
	return searchingWords
}

//標準入力されたワードをスペース区切りで取得し、スライスに格納します
//unitテストを可能にするために引数にio.Readerを用いる
func getSearchingWordsBySpace(stdin io.Reader) []string {
	var searchingWords = []string{}

	sc := bufio.NewScanner(stdin)
	sc.Scan()
	keywords := sc.Text()
	searchingWords = strings.Fields(keywords)

	return searchingWords
}
