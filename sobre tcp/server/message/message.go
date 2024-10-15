package message

import (
	"net"
	"fmt"
)

type M interface {
	SendById(message Message, conexion *net.Conn) bool
}

type Message struct {
	Method string
	Sender string
	Content string
	Receptor string
}


func (m *Message) SendById(conn net.Conn) bool {
	_, err := conn.Write([]byte(m.Sender+ ":" + m.Content))
	if err != nil {
		fmt.Println("Error al escribir:", err)
		return false
	}
	return true
}
