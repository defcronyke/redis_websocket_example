package main

import (
//	"fmt"
//	"log"
	"gopkg.in/redis.v2"
	"time"
	"./RedisUtils"
)

func message_handler(client *redis.Client, msg *redis.Message) {
	if msg.Payload == "rarfle" {
		client.Publish("mychannel", "Ahafghrhshsh!")
	}
}

func main() {
	// Opens a redis connection and subscribes to a channel.
	r_util := RedisUtils.NewRedisUtils(":6379", "mychannel")
	defer r_util.Client.Close()	// We must close the redis connection later.
	defer r_util.Subscribers[0].Close()	// We must close the subscriber later.
	
	for {	// Main loop
		time.Sleep(10 * time.Second)
		pub := r_util.Client.Publish("mychannel", "hello after 10 seconds")
		_ = pub
	}
}