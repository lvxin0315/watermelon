package serv

import (
	"context"
	"github.com/gorilla/websocket"
	wmgrpc "github.com/lvxin0315/watermelon/grpc"
	"google.golang.org/grpc"
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

func (h *Hub) StartGRPC(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := wmgrpc.NewWebSocketServerClient(conn)
	//回去流操作结构体
	stream, err := client.GetData(context.Background())
	if err != nil {
		log.Println("流创建失败,err:", err)
		return
	}
	log.Println("开始读：")
	func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Println("流创建失败,err:", err)
				continue
			}
			h.Broadcast(response.Data)
		}
	}()
}
