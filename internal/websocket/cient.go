package websocket

import (
  "github.com/gorilla/websocket"
  "fmt"
  "errors"
)

type Client struct {
  conn  *websocket.Conn
  id    int
}

func NewClient(conn *websocket.Conn) *Client {
  return &Client{
    conn: conn,
  }
}

func (c *Client) Send(msg string) error {
  if c.conn == nil {
    return errors.New("conn is nil")
  }
  return c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (c *Client) String() string {
  return fmt.Sprintf("Client %d", c.id)
}

type Clients struct {
  clients map[int]*Client
  id int
}

func NewClients() *Clients {
  return &Clients{
    clients: make(map[int]*Client),
    id: 0,
  }
}

func (c *Clients) append(conn *websocket.Conn) *Client {
  client := NewClient(conn)
  c.id++
  client.id = c.id
  c.clients[c.id] = client
  return client
}

func (c *Clients) delete(client *Client) {
  delete(c.clients, client.id)
}

