package websocket

import (
  "github.com/gorilla/websocket"
  "net/http"
  "log"
)

type Message struct {
  msg     []byte
  client  *Client
  msgType int
}

func NewMessage(msg []byte, client *Client) *Message {
  return &Message{
    msg:    msg,
    client: client,
  }
}

type Service struct {
  upgrader  websocket.Upgrader
  clients   *Clients
  msgChan   chan Message
}

func NewService() *Service {
  upgrader := websocket.Upgrader{
    ReadBufferSize:   1024,
    WriteBufferSize:   1024,
  }

  return &Service{
    upgrader: upgrader,
    clients:  NewClients(),
  }
}

func (ws *Service) addClient(conn *websocket.Conn) {
  ws.clients.append(conn)
}

func (ws *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  conn, err := ws.upgrader.Upgrade(w, req, nil)

  if err != nil {
    log.Printf("error during make a new websocket connection !")
    return
  }

  client := ws.clients.append(conn)

  defer func() {
    ws.clients.delete(client)
  }()

  ws.loopClient(client)
}

func (ws *Service) loopClient(client *Client) {
  for {
    msgType, msg, err := conn.ReadMessage()
    if err != nil {
      log.Printf("error during read message")
      return
    }

    ws.msgChan <- NewMessage(msg, client)
  }
}
