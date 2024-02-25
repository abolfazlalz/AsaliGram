package main

import (
  "log"
  "net/http"
  "github.com/abolfazlalz/asaligram/internal/websocket"
)

func main() {
  mux := http.NewServeMux()
  fs := http.FileServer(http.Dir("public"))

  ws := websocket.NewWebsocketService()

  mux.Handle("/ws", ws)
  mux.Handle("/", fs)

  if err := http.ListenAndServe(":8000", mux); err != nil {
    log.Printf("Error during serve and listen: %v", err)
  }
}
