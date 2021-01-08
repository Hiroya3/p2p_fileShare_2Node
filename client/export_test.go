package client

import (
	"io"
)

//clientパッケージとclient_testパッケージのunexportな変数・メソッドをつなげるファイル

func T_getSearchingWordsBySpace(t_stdin io.Reader) []string {
	return getSearchingWordsBySpace(t_stdin)
}
