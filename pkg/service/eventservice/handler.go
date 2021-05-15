package eventservice

import (
	"log"
	"net/http"
	"pishikot/pkg/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type EventService struct {
	Connections map[uint]models.Connection
}

func NewEventService() *EventService {
	return &EventService{
		Connections: map[uint]models.Connection{},
	}
}

func (e *EventService) NewEventConnection(w http.ResponseWriter, r *http.Request, userId uint) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("%s user websocket failed with error %s", userId, err.Error())
		return
	}
	defer ws.Close()
	e.Connections[userId] = models.Connection{
		Conn: ws,
		Send: make(chan interface{}),
	}
	for {
		select {
		case message := <-e.Connections[userId].Send:
			err := ws.WriteJSON(message)
			if err != nil {
				return
			}
		default:
			var _json map[string]interface{}
			err := ws.ReadJSON(&_json)
			if err != nil {
				return
			}
			log.Printf("%s", _json)
		}
	}
}
