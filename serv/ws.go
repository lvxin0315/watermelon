package serv

import (
	"github.com/gorilla/websocket"
	"log"
)

var HubList map[string]*Hub

func init() {
	HubList = make(map[string]*Hub)
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn, send chan []byte) *Client {
	return &Client{hub: hub, conn: conn, send: send}
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//不广播
		//c.hub.broadcast <- message
	}
}

func (c *Client) WritePump() {
	for {
		select {
		case message := <-c.send:
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
		}
	}
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub(uri string) *Hub {
	hub := &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
	HubList[uri] = hub
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Println("len:", len(h.clients))
			for client := range h.clients {
				select {
				case client.send <- message:
					//log.Println(message)

				default:
					log.Println("close")
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) Broadcast(m []byte) {
	h.broadcast <- m
}
