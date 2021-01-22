package p2p

import "net"

func T_readRequestMessage(conn net.Conn) ([]string, error) {
	messageSlice, err := readRequestMessage(conn)
	return messageSlice, err
}

func T_compareHash(requestBodyStr, requestHash string) bool {
	return compareHash(requestBodyStr, requestHash)
}
