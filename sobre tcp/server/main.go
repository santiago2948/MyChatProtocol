package main

import (
	"fmt"
	"net"
	"server/socket"
	"server/client"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al crear el socket:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP escuchando en :8080")
	var clients = make(map[string]*client.Client)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		ws := socket.Socket{Clients: clients}

		go ws.HandleConnection(conn)
	}
}
