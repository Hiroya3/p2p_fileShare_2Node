package service

import (
	"io"
)

func T_getSearchingWordsBySpace(t_stdin io.Reader) []string {
	return getSearchingWordsBySpace(t_stdin)
}
