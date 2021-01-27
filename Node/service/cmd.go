package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//コマンドラインに検索を表示します。
func GetSearchingWords() []string {
	searchingWords := []string{}
	for {
		fmt.Println("ファイルを検索します。キーワードをスペース区切りで3つまで指定してください。\n３つ以上指定した場合ははじめ３つが採用されます。")
		searchingWords = getSearchingWordsBySpace(os.Stdin)

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

//相手ノードから検索したファイルの検索一覧の中からダウンロードしたいファイルを検索する
func SelectFileName(files []string) string {
	//検索候補がない場合はlen(files):0 で files[0]:"" となる配列が返る
	if files[0] == "" {
		fmt.Println("候補がありませんでした。\n時間をおいて再度検索して下さい。")
		os.Exit(0)
	}
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("下記からダウンロードしたいファイルの番号を選択して下さい。\nない場合は0を入力して終了。")
		for i := 0; i < len(files); i++ {
			fmt.Printf(strconv.Itoa(i+1) + " :" + files[i] + "\n")
		}
		fmt.Print("ファイル番号 >")
		sc.Scan()
		fileNumStr := sc.Text()
		if fileNumStr == "0" {
			fmt.Println("また使って下さいね！")
			os.Exit(0)
		}

		fileNum, err := strconv.Atoi(fileNumStr)
		if err != nil {
			log.Println(err)
		}

		if fileNum <= len(files) && fileNum > 0 {
			return files[fileNum-1]
		}

		fmt.Println("選択肢以外が入力されました。再度入力して下さい。")
	}
}
