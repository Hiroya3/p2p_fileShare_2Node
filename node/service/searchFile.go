package service

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//files/upload内にあるファイルを検索
const targetPath = "../files/upload"

// searchingWords : 相手ノードからきた検索ワード（最大3つ）
func SearchLocalFiles(seachingWords []string) []string {
	filNameSlice := []string{}

	//path ディレクトリ内の1つ1つのファイルへのpath
	//info アクセスしているファイルの情報
	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("%qへの検索の際にエラーが発生しました。エラー：%s", path, err)
		}
		if info.IsDir() {
			return nil
		}

		for i := 0; i < len(seachingWords); i++ {
			if strings.Contains(info.Name(), seachingWords[i]) {
				filNameSlice = append(filNameSlice, info.Name())
			}
		}
		return nil
	})
	if err != nil {
		log.Println("指定されたパスへの検索でエラーが発生しました。")
	}
	return filNameSlice
}
