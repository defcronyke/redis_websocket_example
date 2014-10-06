package RedisUtils

import(
//	"fmt"
	"log"
	"gopkg.in/redis.v2"
)

type redisUtils struct {
	Client *redis.Client
	Subscribers []*redis.PubSub
}

// Message handlers need this signature.
type MessageHandler func(*redis.Client, *redis.PubSub, *redis.Message)

// Constructor
func NewRedisUtils(url, channel string) *redisUtils {
	r := new(redisUtils)
	r.Client = r.ConnectToRedis(url)
	subscriber := r.Subscribe(r.Client, channel)
	r.Subscribers = append(r.Subscribers, subscriber)
	r.ListenToSubscription(r.Client, subscriber, r.MessageHandlerDefault)
	
	return r
}

func (r redisUtils) ConnectToRedis(url string) *redis.Client {
	log.Printf("Connecting to: %v\n", url)
	client := redis.NewTCPClient(&redis.Options{
		Addr: url,
	})
	client.FlushDb()
	return client
}

func (r redisUtils) Subscribe(client *redis.Client, channel string) *redis.PubSub {
	sub := client.PubSub()
	err := sub.Subscribe(channel)
	if err != nil {
		log.Printf("Error: %v", err)
		return sub
	}
	log.Printf("Subscribed to channel: %v", channel)
	
	return sub
}

func (r redisUtils) ListenToSubscription(client *redis.Client, subscriber *redis.PubSub, message_handler MessageHandler) {
	go func(subscriber *redis.PubSub) { // listen for messages in a goroutine
		for {
			msg, err := subscriber.Receive()
			if err != nil {
				log.Printf("Error: %v", err)
			}
		
			switch t := msg.(type) {
		
			case *redis.Message:
				log.Printf("%v: %v\n", t.Channel, t.Payload)
				message_handler(client, subscriber, t)
		
			default:
				log.Printf("%v\n", t)
			}
		}
	}(subscriber)
}

func (r redisUtils) MessageHandlerDefault(client *redis.Client, subscriber *redis.PubSub, message *redis.Message) {
	if message.Payload == "hi" {
		log.Printf("%v: bye\n", message.Channel)
	}
}
