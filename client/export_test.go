package client

import (
	"io"
)

//clientパッケージとclient_testパッケージのunexportな変数・メソッドをつなげるファイル

func T_GetSearchingWords(t_stdin io.Reader) []string {
	return getSearchingWords(t_stdin)
}
