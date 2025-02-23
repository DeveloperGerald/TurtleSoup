package deepseek

import (
	"log"

	"github.com/DeveloperGerald/TurtleSoup/config"
)

type Client struct {
	SecretKey string
	Domain    string
}

var client *Client

func Init() {
	config := config.GetConfig().Deepseek
	client = &Client{
		SecretKey: config.SecretKey,
		Domain:    config.Domain,
	}
	log.Println("deepseek client init success")
}

func GetClient() *Client {
	return client
}
