package WebSocketConn

import (
	"net/http"
	"log"
)

type WebSocketHandlerT func(http.ResponseWriter, *http.Request) // websocket handlers need this signature

type webSocketConn struct {
}

func NewWebSocketConn() *webSocketConn {
	w := new(webSocketConn)
	return w
}

func (w *webSocketConn) Serve(url string, websocket_handler WebSocketHandlerT) {
	http.Handle("/", http.FileServer(http.Dir("./webroot")))
	http.HandleFunc("/ws", websocket_handler)
	log.Printf("Listening on: http://%v", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
