package main

import (
	"fmt"
	"net"
)

func main() {
	// Crear socket y vincular a una dirección
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al crear el socket:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP escuchando en :8080")

	for {
		// Aceptar conexiones entrantes
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}

		// Manejar la conexión en una goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Leer datos del cliente
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer:", err)
			return
		}

		mensaje := string(buffer[:n])
		fmt.Printf("Recibido: %s\n", mensaje)

		// Enviar respuesta al cliente
		_, err = conn.Write([]byte("Mensaje recibido: " + mensaje))
		if err != nil {
			fmt.Println("Error al escribir:", err)
			return
		}
	}
}