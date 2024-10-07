package message

type Mensaje struct {
	SenderID string `json:"sender_id"`
	ReceptorId string `json:"receptor_id"`
	Content string `json:"content"`
}
