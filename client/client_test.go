package client_test

import (
	"bytes"
	"p2p_fileShare_2Node/client"
	"testing"
)

func TestGetSearchingWords(t *testing.T) {
	//参考：https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
	var stdin_0word, stdin_1word, stdin_3words, stdin_4words bytes.Buffer

	expect_1word := []string{"aaa"}
	expect_3words := []string{"aaa", "bbb", "ccc"}
	expect_4words := []string{"aaa", "bbb", "ccc"}

	stdin_0word.Write([]byte("\n"))
	stdin_1word.Write([]byte("aaa\n"))
	stdin_3words.Write([]byte("aaa bbb ccc\n"))
	stdin_4words.Write([]byte("aaa bbb ccc ddd\n"))

	result_0word := client.T_GetSearchingWords(&stdin_0word)
	result_1word := client.T_GetSearchingWords(&stdin_1word)
	result_3words := client.T_GetSearchingWords(&stdin_3words)
	result_4words := client.T_GetSearchingWords(&stdin_4words)

	if len(result_0word) != 0 {
		t.Errorf("length is more than 0. got:%s", result_0word)
	}

	for i := 0; i < len(result_1word); i++ {
		if expect_1word[i] != result_1word[i] {
			t.Errorf("got: %s, expect_3wordsed: %s", result_1word[i], expect_1word[i])
		}
	}

	for i := 0; i < len(result_3words); i++ {
		if expect_3words[i] != result_3words[i] {
			t.Errorf("got: %s, expect_3wordsed: %s", result_3words[i], expect_3words[i])
		}
	}

	for i := 0; i < len(result_4words); i++ {
		if i < 3 {
			if expect_4words[i] != result_4words[i] {
				t.Errorf("got: %s, expect_3wordsed: %s", result_4words[i], expect_4words[i])
			}
		} else {
			//３つまでしか取得できないので超えた分は無視
			break
		}
	}
}
