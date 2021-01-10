package server

import "net"

func T_readRequestMessage(conn net.Conn) ([]byte, int) {
	messageBuff, messageLen := readRequestMessage(conn)
	return messageBuff, messageLen
}
