package main

import (
	"strconv"
)
// Hub maintains the set of active clients and broadcasts messages to the
// clients.
var id = 1

type Hub struct {
	// Registered clients.
	clients map[*Client]ClientStruct

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type ClientStruct struct {
	connected bool

	id int
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]ClientStruct),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = ClientStruct {
				connected: true,
				id: id,
			} 
			client.send <- []byte(strconv.Itoa(id)) 
			id++
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}