// demo from https://www.youtube.com/watch?v=JuUAEYLkGbM&t=415s
package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"time"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		fmt.Println("received msg ", string(msg))

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write err", err)
			}
		}(ws)
	}
}

func (s *Server) handleWsOrderbook(ws *websocket.Conn) {

	fmt.Println("new incoming connection from client to orderbook feed", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data -> %d \n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbook", websocket.Handler(server.handleWsOrderbook))
	http.ListenAndServe(":3000", nil)
	fmt.Println("serer is inited")
}
