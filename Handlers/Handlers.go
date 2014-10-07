package Handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"gopkg.in/redis.v2"
	"github.com/gorilla/websocket"
)

type RedisClientT struct {
	RedisClient *redis.Client
	RedisSubscribers []*redis.PubSub
}

type RedisClientsT []RedisClientT
type WebSocketClientsT []*websocket.Conn
type RedisHandlerT func(*redis.Message)
type WebSocketHandlerT func(http.ResponseWriter, *http.Request)

type handlers struct {
	RedisClients RedisClientsT
	WebSocketClients WebSocketClientsT
	Upgrader websocket.Upgrader
}

func NewHandlers() *handlers {
	h := new(handlers)
	h.Upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	return h
}

// An example redis handler.
func (h handlers) RedisHandlerDefault(redis_msg *redis.Message) {
	if redis_msg.Payload == "hi" {
		log.Printf("%v: bye\n", redis_msg.Channel)
	}
}

// An example websocket handler.
func (h handlers) WebSocketHandlerDefault(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", 405)
		return
	}
	
	ws, err := h.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	
	go func() {	// Read incoming messages.
		for {
			msg_type, msg, err := ws.NextReader()
			if err != nil {	// If page closed.
				return
			}
			_ = msg_type
			msg_bytes, err := ioutil.ReadAll(msg)
			if err != nil {
				log.Printf("Error: %v", err)
				return
			}
			log.Printf("Incoming WebSocket msg: %v", string(msg_bytes))
		}
	}()
}