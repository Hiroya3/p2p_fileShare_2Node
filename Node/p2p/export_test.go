package p2p

import "net"

func TreadRequestMessage(conn net.Conn) ([]string, error) {
	messageSlice, err := readRequestMessage(conn)
	return messageSlice, err
}

func TcompareHash(requestBodyStr, requestHash string) bool {
	return compareHash(requestBodyStr, requestHash)
}
