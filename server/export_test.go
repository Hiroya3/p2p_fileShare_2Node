package server

import "net"

func T_readRequestMessage(conn net.Conn) []string {
	messageSlice := readRequestMessage(conn)
	return messageSlice
}
