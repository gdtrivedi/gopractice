package timestamp

import (
	"context"
	"fmt"
	"time"
)

// BackoffAndDeadLetteringPubSubClientConfig configuration of Backoff And Deadlettering PubSub Client
type BackoffAndDeadLetteringPubSubClientConfig struct {
	WaitForDlt *time.Duration
}

// BackoffAndDeadLetteringPubSubClient wraps another PubSubClient and has backsoff and deadletter behavior when Nack is called
type BackoffAndDeadLetteringPubSubClient struct {
	ctx    context.Context
	config *BackoffAndDeadLetteringPubSubClientConfig
}

func defaultConfig() *BackoffAndDeadLetteringPubSubClientConfig {
	defaultWaitForDlt := 1 * time.Hour

	return &BackoffAndDeadLetteringPubSubClientConfig{
		WaitForDlt: &defaultWaitForDlt,
	}
}

// NewBackoffAndDeadLetteringPubSubClient creates an instance of BackoffAndDeadLetteringPubSubClient
func NewBackoffAndDeadLetteringPubSubClient(ctx context.Context,
	config *BackoffAndDeadLetteringPubSubClientConfig) *BackoffAndDeadLetteringPubSubClient {
	client := &BackoffAndDeadLetteringPubSubClient{
		ctx: ctx,
	}

	client.config = config

	// Set default config.
	if client.config == nil {
		client.config = defaultConfig()
	}

	return client
}

func TimeInitTest() {
	waitForDlt := 24 * time.Hour
	config := BackoffAndDeadLetteringPubSubClientConfig{WaitForDlt: &waitForDlt}

	client := NewBackoffAndDeadLetteringPubSubClient(context.Background(), &config)
	fmt.Printf(client.config.WaitForDlt.String())
}
