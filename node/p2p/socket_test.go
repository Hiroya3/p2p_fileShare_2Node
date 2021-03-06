package p2p_test

import (
	"p2p_fileShare_2Node/node/p2p"
	"testing"
)

func TestCompareHash(t *testing.T) {

	//リクエストの文字列
	requestBodyStr := "001:searchWords:aa,bb,cc:"
	//ハッシュ値
	hash := "07c53b87cdf1f44b865d88254472a5be50bbdb3de41cb56c2b5f5fb1b50747fc"
	if !p2p.TcompareHash(requestBodyStr, hash) {
		t.Errorf("%sのハッシュ値（%s）の生後判定が誤っています。", requestBodyStr, hash)
	}
}

func TestCreateRequestStr(t *testing.T) {
	expectedStr := "001:searchWords:aa,bb,cc:07c53b87cdf1f44b865d88254472a5be50bbdb3de41cb56c2b5f5fb1b50747fc"

	bodySlice := []string{"aa,bb,cc"}
	if expectedStr != p2p.TcreateRequestStr("001", "searchWords", bodySlice) {
		t.Errorf("expectedResult : %s\nbut\nacutual Result : %s", expectedStr, p2p.TcreateRequestStr("001", "searchWords", bodySlice))
	}
}
