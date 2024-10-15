package socket

import (
	"server/message"
	"server/client"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
)

type SocketInterface interface{
	HandleConnections(w http.ResponseWriter, r *http.Request)
	Conexion()
	NewClient()
	ExitClient(clientId string)
	HandleMessage(ws *websocket.Conn, clienteID string)
	
}
type Socket struct {
	Clients map[string]*client.Client
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func (s *Socket) Conexion(){
	http.HandleFunc("/", s.HandleConnections)

	fmt.Println("Servidor WebSocket escuchando en :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}

func (s *Socket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Println("Error leyendo mensaje inicial:", err)
		msg := message.Mensaje{Type: "rejectConexion", SenderID: "", ReceptorId: "", Content: ""}
		rejectMessage, err := json.Marshal(msg)
		if err!=nil{

		}
		ws.WriteMessage(websocket.TextMessage, rejectMessage)
		return
	}

	var mensajeInicial message.Mensaje
	err = json.Unmarshal(msg, &mensajeInicial)
	if err != nil {
		log.Println("Error al decodificar mensaje JSON inicial:", err)
		return
	}

	clienteID := mensajeInicial.SenderID

	if s.Clients[clienteID] != nil {
		fmt.Printf("Cliente ya conectado con ID: %s\n", clienteID)
		return
	}else{
		fmt.Printf("Nuevo cliente conectado con ID: %s\n", clienteID)
		nuevoCliente := &client.Client{Conn: ws, Id: clienteID}
		s.Clients[clienteID] = nuevoCliente
		s.NewClient(clienteID)
	}
	s.HandleMessage(ws, clienteID)

}

func (s *Socket) HandleMessage(ws *websocket.Conn, clienteID string){
	for {
		// Leer los siguientes mensajes del cliente
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error leyendo mensaje:", err)
			delete(s.Clients, clienteID)
			s.ExitClient(clienteID)
			break
		}
		
		var mensaje message.Mensaje
		err = json.Unmarshal(msg, &mensaje)
		if err != nil {
			log.Println("Error al decodificar mensaje JSON:", err)
			return
		}
		receptor := s.Clients[mensaje.ReceptorId]
		if receptor != nil {
		err = receptor.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error enviando mensaje JSON:", err)
			break
		}
	}else{
		fmt.Println("No se encontró el receptor para el mensaje")
	}
}
}

func (s *Socket) ExitClient(clientId string){
	for client, conn := range s.Clients{
		msg:= message.Mensaje{Type: "exitClient", SenderID: clientId, ReceptorId: client, Content: ""}
		msgc, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			return
		}else  {
			conn.Conn.WriteMessage(websocket.TextMessage, msgc)
		}
	}
}


func (s *Socket) NewClient(id string){

	currentConn := s.Clients[id]

	for client, conn := range s.Clients{
		msg:= message.Mensaje{Type: "newClient", SenderID: id, ReceptorId: client, Content: ""}
		msgToCurrent:= message.Mensaje{Type: "newClient", SenderID: client, ReceptorId: id, Content: ""}
		msgc, err1 := json.Marshal(msg)
		msgcToCurrent, err2 := json.Marshal(msgToCurrent)
		if err1 != nil || err2 != nil {
			fmt.Println("Error al convertir a JSON:", err1)
			return
		}else if id!= client {
			currentConn.Conn.WriteMessage(websocket.TextMessage, msgcToCurrent)
			conn.Conn.WriteMessage(websocket.TextMessage, msgc)
		}
	}
}