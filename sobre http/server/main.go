package main

import (
	"encoding/json"
	"server/message"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


type Cliente struct {
	conn *websocket.Conn
	id   string
}

var clientes = make(map[string]*Cliente)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Establecer la conexión WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Leer el primer mensaje que contiene el ClienteID
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Println("Error leyendo mensaje inicial:", err)
		return
	}

	// Deserializar el mensaje para obtener el ClienteID
	var mensajeInicial message.Mensaje
	err = json.Unmarshal(msg, &mensajeInicial)
	if err != nil {
		log.Println("Error al decodificar mensaje JSON inicial:", err)
		return
	}

	clienteID := mensajeInicial.SenderID
	if clientes[clienteID] != nil {
		fmt.Printf("Cliente ya conectado con ID: %s\n", clienteID)
		return
	}else{
		fmt.Printf("Nuevo cliente conectado con ID: %s\n", clienteID)

	// Crear un nuevo cliente y agregarlo al mapa
		nuevoCliente := &Cliente{conn: ws, id: clienteID}
		clientes[clienteID] = nuevoCliente

	}
	
	for {
		// Leer los siguientes mensajes del cliente
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error leyendo mensaje:", err)
			delete(clientes, clienteID)
			break
		}

		fmt.Printf("Mensaje recibido del cliente %s: %s\n", clienteID, string(msg))
		//mensaje2 := json.Unmarshal(msg, )
		var mensaje message.Mensaje
		err = json.Unmarshal(msg, &mensaje)
		if err != nil {
			log.Println("Error al decodificar mensaje JSON inicial:", err)
			return
		}
		receptor := clientes[mensaje.ReceptorId]
		fmt.Println("Receptor ID:",mensaje.ReceptorId)
		if receptor != nil {
		err = receptor.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error enviando mensaje JSON:", err)
			break
		}
	}else{
		fmt.Println("No se encontró el receptor para el mensaje")
	}
}
}

func main() {
	http.HandleFunc("/", handleConnections)

	fmt.Println("Servidor WebSocket escuchando en :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
