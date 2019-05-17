package pubsub

import (
	"context"
	"log"
	"os"
)

var projectID = os.Getenv("GCP_PROJECT")

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	log.Printf("[init] project: %v", projectID)
}

func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s", name)
	return nil
}
