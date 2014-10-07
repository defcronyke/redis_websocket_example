package main

import (
//	"fmt"
//	"log"
//	"gopkg.in/redis.v2"
//	"time"
	"./RedisConn"
	"./WebSocketConn"
	"./Handlers"
)

func main() {	
	handlers := Handlers.NewHandlers()	// Handlers for redis and websocket messages.
	ws_conn := WebSocketConn.NewWebSocketConn()	// The websocket connection.
	r_conn := RedisConn.NewRedisConn(":6379", "mychannel")	// The redis connection.
	defer r_conn.RedisClient.Close()	// We must close the redis connection later.
	defer r_conn.RedisSubscribers[0].Close()	// We must close the redis subscriber later.
	
	redis_client := Handlers.RedisClientT{ r_conn.RedisClient, r_conn.RedisSubscribers } // Package up redis connection info for the handlers.
	handlers.RedisClients = append(handlers.RedisClients, redis_client)	// Add redis client to the handlers object.
	
	r_conn.ListenToSubscription(handlers.RedisHandlerDefault)	// Listen for redis messages.
	ws_conn.Serve("localhost:8080", handlers.WebSocketHandlerDefault)	// Listen for websocket messages.
	
/*  // Uncomment this if only using redis, and not the web server.
	for {
		time.Sleep(10 * time.Second)
		pub := r_conn.RedisClient.Publish("mychannel", "hello after 10 seconds")
		_ = pub
	}
*/
}