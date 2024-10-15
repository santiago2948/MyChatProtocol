package socket

import (
	"server/client"
	"server/message"
	"net"
	"fmt"
	"strings"
)

type SocketInterfface interface{
	HandleConnection(conn net.Conn)
	Conexions()
}


type Socket struct{
	Clients map[string]*client.Client
	Listener net.Listener
}

func (s *Socket) Conexion(){
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}

		go s.HandleConnection(conn)
	}
}


func (s *Socket) HandleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			for nickname, client := range s.Clients {
				if client.Conexion == conn {
					delete(s.Clients, nickname)
					fmt.Printf("Cliente %s desconectado\n", nickname)
					break
				}
			}
			return
		}

		mensaje := string(buffer[:n])
		msg := strings.Split(mensaje, "field//")

		if msg[1] == "connect" {
			fmt.Println("nuevo cliente conectado", msg[2])
			newClient := &client.Client{Conexion: conn, Nickname: msg[2]}
			s.Clients[msg[2]] = newClient
		}else if len(msg) > 2 {
			
		message := message.Message{Method: msg[1], Sender: msg[2], Content: msg[4], Receptor: msg[3]}
			connReceptor := s.Clients[message.Receptor]
			if connReceptor != nil {
			response:= message.SendById(connReceptor.Conexion)
			if !response {fmt.Println("Error al enviar el mensaje")}
			}else{
				fmt.Println("no se pudo encontrar al receptor", message.Receptor)
			}
			
		}
}
}