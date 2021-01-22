package p2p

import "net"

func T_readRequestMessage(conn net.Conn) []string {
	messageSlice := readRequestMessage(conn)
	return messageSlice
}

func T_compareHash(requestBodyStr, requestHash string) bool {
	return compareHash(requestBodyStr, requestHash)
}
