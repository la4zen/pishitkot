package models

import (
	"github.com/gorilla/websocket"
)

type CodeEvent struct {
	FromUser uint
	TaskID   uint
	Code     string
}

type CreateEvent struct {
	FromUser  uint
	FirstName string
	LastName  string
	TaskId    uint
}

type Connection struct {
	Conn *websocket.Conn
	Send chan interface{}
}
