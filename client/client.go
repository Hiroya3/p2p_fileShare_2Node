package client

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type SearchingWords struct {
	Stdin    io.Reader
	Keyword1 string
	Keyword2 string
	Keyword3 string
}

func createSearchingWords() *SearchingWords {
	return &SearchingWords{
		Stdin:    os.Stdin,
		Keyword1: "",
		Keyword2: "",
		Keyword3: "",
	}
}

//コマンドラインに案内画面を表示します。
func ViewCmd() {
	SearchingWords := createSearchingWords()
	for {
		fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。\n３つ以上指定した場合ははじめ３つが採用されます。")
		words := getWords()

		if len(words) > 3 {
			fmt.Println("4つ以上指定されたため、はじめの３つを取得します。")
		}
		if len(words) > 0 {
			SearchingWords.Keyword1 = words[0]
			SearchingWords.Keyword2 = words[1]
			SearchingWords.Keyword3 = words[2]
			break
		}
		if len(words) == 0 {
			fmt.Println("キーワードを指定してください。")
		}
	}
}

//検索ワードを3つ取得します。4つ以上入力された場合ははじめの3つが取得されます。
func getWords() []string {
	var searchingWords = []string{}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	keywords := sc.Text()
	searchingWords = strings.Fields(keywords)

	return searchingWords
}
