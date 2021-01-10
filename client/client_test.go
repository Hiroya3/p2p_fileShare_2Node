package client_test

import (
	"bytes"
	"p2p_fileShare_2Node/client"
	"testing"
)

func TestGetSearchingWordsBySpace(t *testing.T) {
	//参考：https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
	var stdin_non_word, stdin_words bytes.Buffer

	expect_words := []string{"aaa", "bbb", "ccc", "ddd"}

	stdin_non_word.Write([]byte("\n"))
	stdin_words.Write([]byte(" aaa  bbb　ccc	ddd   \n"))

	result_non_word := client.T_getSearchingWordsBySpace(&stdin_non_word) //byteのポインター型にするのはos.Stdinがfileのポインター型であり型を揃えるため
	result_words := client.T_getSearchingWordsBySpace(&stdin_words)

	if len(result_non_word) != 0 {
		t.Errorf("length is more than 0. got:%s", result_non_word)
	}

	for i := 0; i < len(result_words); i++ {
		if expect_words[i] != result_words[i] {
			t.Errorf("got: %s, expect_words: %s", result_words[i], expect_words[i])
		}
	}
}
