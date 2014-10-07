package RedisConn

import(
//	"fmt"
	"log"
	"gopkg.in/redis.v2"
)

type redisConn struct {
	RedisClient *redis.Client
	RedisSubscribers []*redis.PubSub
}

type RedisHandlerT func(*redis.Message) // Redis handlers need this signature.

// Constructor
func NewRedisConn(url, channel string) *redisConn {
	r := new(redisConn)
	r.RedisClient = r.ConnectToRedis(url)
	redis_subscriber := r.Subscribe(r.RedisClient, channel)
	r.RedisSubscribers = append(r.RedisSubscribers, redis_subscriber)
	return r
}

func (r redisConn) ConnectToRedis(url string) *redis.Client {
	log.Printf("Connecting to: %v\n", url)
	client := redis.NewTCPClient(&redis.Options{
		Addr: url,
	})
	client.FlushDb()
	return client
}

func (r redisConn) Subscribe(client *redis.Client, channel string) *redis.PubSub {
	sub := client.PubSub()
	err := sub.Subscribe(channel)
	if err != nil {
		log.Printf("Error: %v", err)
		return sub
	}
	log.Printf("Subscribed to channel: %v", channel)
	
	return sub
}

func (r redisConn) ListenToSubscription(redis_handler RedisHandlerT) {
	go func(subscriber *redis.PubSub) { // listen for messages in a goroutine
		for {
			msg, err := subscriber.Receive()
			if err != nil {
				log.Printf("Error: %v", err)
			}
		
			switch t := msg.(type) {
		
			case *redis.Message:
				log.Printf("%v: %v\n", t.Channel, t.Payload)
				redis_handler(t)
		
			default:
				log.Printf("%v\n", t)
			}
		}
	}(r.RedisSubscribers[0])	// TODO: make this work for all subscribers.
}