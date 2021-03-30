package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	staticFileHandler := http.FileServer(http.Dir("src/herry2038.org/websocket/static"))
	var siteMux = http.NewServeMux()
	siteMux.Handle("/", http.StripPrefix("/", staticFileHandler))
	http.Handle("/", http.Handler(siteMux))
	http.Handle("/ws", generateHandleWS())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
func generateHandleWS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("New client connected: %s", r.RemoteAddr)
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}
		upgrader := &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    []string{"test"},
			CheckOrigin:     nil,
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("closed %v", err)
			return
		}
		defer conn.Close()
		processWSConn(conn)
	}
}
func processWSConn(conn *websocket.Conn) {
	for {
		typ, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("receive message failed:%v", err)
			return
		}
		if typ != websocket.TextMessage {
			continue
		}
		fmt.Printf("received a msg:%s\n", string(msg))
		backMsg := "hello:" + string(msg)
		conn.WriteMessage(websocket.TextMessage, []byte(backMsg))
	}
}
