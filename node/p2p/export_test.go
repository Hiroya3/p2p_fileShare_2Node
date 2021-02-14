package p2p

import "net"

func TreadRequestMessage(conn net.Conn) ([]string, error) {
	_, messageSlice, err := readRequestMessage(conn)
	return messageSlice, err
}

func TcompareHash(requestBodyStr, requestHash string) bool {
	return compareHash(requestBodyStr, requestHash)
}

func TcreateRequestStr(headerNumStr, methodStr string, bodySlice []string) string {
	return createRequestStr(headerNumStr, methodStr, bodySlice)
}
