package WebSocketConn

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

type WebSocketHandlerT func(http.ResponseWriter, *http.Request) // websocket handlers need this signature

type webSocketConn struct {
	Upgrader websocket.Upgrader
}

func NewWebSocketConn() *webSocketConn {
	w := new(webSocketConn)
	w.Upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	return w
}

func (w webSocketConn) Serve(url string, websocket_handler WebSocketHandlerT) {
	http.Handle("/", http.FileServer(http.Dir("./webroot")))
	http.HandleFunc("/ws", websocket_handler)
	log.Printf("Listening on: http://%v", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
