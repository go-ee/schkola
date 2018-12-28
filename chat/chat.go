package chat

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

type Client struct {
	Id      string
	Socket  *websocket.Conn
	Send    chan []byte
	manager *ClientManager
}

func (o *ClientManager) NewClient(conn *websocket.Conn) *Client {
	return &Client{Id: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte), manager: o}
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

func (o *ClientManager) Start() {
	for {
		select {
		case conn := <-o.Register:
			o.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new Socket has connected."})
			o.Send(jsonMessage, conn)
		case conn := <-o.Unregister:
			if _, ok := o.Clients[conn]; ok {
				close(conn.Send)
				delete(o.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A Socket has disconnected."})
				o.Send(jsonMessage, conn)
			}
		case message := <-o.Broadcast:
			for conn := range o.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(o.Clients, conn)
				}
			}
		}
	}
}

func (o *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range o.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (c *Client) Read() {
	defer func() {
		c.manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			c.manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.Id, Content: string(message)})
		c.manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (o *ClientManager) Handler(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := o.NewClient(conn)

	o.Register <- client

	go client.Read()
	go client.Write()
}
