package client_test

import (
	"bytes"
	"p2p_fileShare_2Node/client"
	"testing"
)

func TestGetSearchingWords(t *testing.T) {
	var stdin bytes.Buffer

	expect := []string{"aaa", "bbb", "ccc"}
	stdin.Write([]byte("aaa bbb ccc\n"))

	result := client.T_GetSearchingWords(&stdin)

	for i := 0; i < len(expect); i++ {
		if expect[i] != result[i] {
			t.Errorf("got: %s, expected: %s", result[i], expect[i])
		}
	}
}
