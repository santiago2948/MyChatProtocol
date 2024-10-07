package main

import (
	"fmt"
	"net"
	"strings"
	"server/message"
)


type Client struct {
	conexion net.Conn
	Nickname string
}

var clients = make(map[string]*Client)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al crear el socket:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP escuchando en :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			for nickname, client := range clients {
				if client.conexion == conn {
					delete(clients, nickname)
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
			newClient := &Client{conexion: conn, Nickname: msg[2]}
			clients[msg[2]] = newClient
		}else if len(msg) > 2 {
			
		message := message.Message{Method: msg[1], Sender: msg[2], Content: msg[4], Receptor: msg[3]}
			connReceptor := clients[message.Receptor]
			if connReceptor != nil {
			response:= message.SendById(connReceptor.conexion)
			if !response {fmt.Println("Error al enviar el mensaje")}
			}else{
				fmt.Println("no se pudo encontrar al receptor", message.Receptor)
			}
			
		}
}
}