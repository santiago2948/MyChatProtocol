package main

import (
	"server/client"
	"server/socket"
)

func main() {
	var clientes = make(map[string]*client.Client)
	ws := socket.Socket{Clients: clientes}
	ws.Conexion()
}


