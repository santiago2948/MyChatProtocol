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
	ws := socket.Socket{Clients: clients, Listener: listener}
	ws.Conexion()
}
