package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

// to use web sockets, we must upgrade the HTTP connection using the
// websocket.Upgrader type, which is reusable so we need only create one.
var upgrader = &websocket.Upgrader{
	ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte

	// The join and leave channels exist simply to allow
	// us to safely add and remove clients from the clients map.
	// If we were to access the map directly, it is possible that two Go routines running
	// concurrently might try to modify the map at the same time resulting in corrupt
	// memory or an unpredictable state.
	join chan *client  // join is a channel for clients wishing to join the room.
	leave chan *client // leave is a channel for clients wishing to leave the room.

	// clients holds all current clients in this room.
	clients map[*client]bool
}

// newRoom makes a new room that is ready to go.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// we get the socket
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	// we then create our client and pass
	// it into the join channel for the current room.
	client := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}

	r.join <- client
	defer func() { r.leave <- client }()

	go client.write()

	// block operations (keeping the connection alive)
	// until it's time to close it
	client.read()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true

		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)

		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
