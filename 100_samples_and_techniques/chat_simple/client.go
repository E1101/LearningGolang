package main

import (
	"github.com/gorilla/websocket"
)

// Client:
// client represents a single chatting user.

type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan []byte
	// room is the room this client is chatting in.
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			// If it encounters an error will break and
			// the socket will be closed.
			break
		}
	}

	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			// If writing to the socket fails, the for loop is broken
			// and the socket is closed
			break
		}
	}

	c.socket.Close()
}
