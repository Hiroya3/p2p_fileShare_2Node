package client

import (
	"bufio"
	"os"
)

//コマンドラインに表示する検索画面の初期表示
func GetSearchingWords() []string {
	var searchingWords = []string{}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		word := sc.Text()
		searchingWords = append(searchingWords, word)
	}
	return searchingWords
}
