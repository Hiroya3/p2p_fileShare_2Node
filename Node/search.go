package server

import (
	"log"
	"os"
	"path/filepath"
)

//files/upload内にあるファイルを検索

const targetPath = "../files/upload/"

// searchingWords : 相手ノードからきた検索ワード（最大3つ）
func SearchLocalFiles(seachingWords []string) []string {
	//path ディレクトリ内の1つ1つのファイルへのpath
	//info アクセスしているファイルの情報
	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("%qへの検索の際にエラーが発生しました。エラー：%s", path, err)
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		//TODO:ファイル検索のロジックの実装必要あり
		return nil
	})
	if err != nil {
		log.Fatalln("指定されたパスへの検索でエラーが発生しました。")
	}
	return nil
}
