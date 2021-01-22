package p2p_test

import (
	"p2p_fileShare_2Node/Node/p2p"
	"testing"
)

func TestCompareHash(t *testing.T) {

	//リクエストの文字列
	requestBodyStr := "001:searchWords:aa,bb,cc:"
	//ハッシュ値
	hash := "07c53b87cdf1f44b865d88254472a5be50bbdb3de41cb56c2b5f5fb1b50747fc"
	if !p2p.T_compareHash(requestBodyStr, hash) {
		t.Errorf("%sのハッシュ値（%s）の生後判定が誤っています。", requestBodyStr, hash)
	}
}
