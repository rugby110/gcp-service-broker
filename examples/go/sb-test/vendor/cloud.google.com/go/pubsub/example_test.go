// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub_test

import (
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	_, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	// See the other examples to learn how to use the Client.
}

func ExampleClient_CreateTopic() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	// Create a new topic with the given name.
	topic, err := client.CreateTopic(ctx, "topicName")
	if err != nil {
		// TODO: Handle error.
	}

	_ = topic // TODO: use the topic.
}

func ExampleClient_CreateSubscription() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	// Create a new topic with the given name.
	topic, err := client.CreateTopic(ctx, "topicName")
	if err != nil {
		// TODO: Handle error.
	}

	// Create a new subscription to the previously created topic
	// with the given name.
	sub, err := client.CreateSubscription(ctx, "subName", topic, 10*time.Second, nil)
	if err != nil {
		// TODO: Handle error.
	}

	_ = sub // TODO: use the subscription.
}

func ExampleTopic_Delete() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	topic := client.Topic("topicName")
	if err := topic.Delete(ctx); err != nil {
		// TODO: Handle error.
	}
}

func ExampleTopic_Exists() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	topic := client.Topic("topicName")
	ok, err := topic.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	if !ok {
		// Topic doesn't exist.
	}
}

func ExampleTopic_Publish() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	topic := client.Topic("topicName")
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}

func ExampleTopic_Subscriptions() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	topic := client.Topic("topic-name")
	// List all subscriptions of the topic (maybe of multiple projects).
	for subs := topic.Subscriptions(ctx); ; {
		sub, err := subs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		_ = sub // TODO: use the subscription.
	}
}

func ExampleSubscription_Delete() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	sub := client.Subscription("subName")
	if err := sub.Delete(ctx); err != nil {
		// TODO: Handle error.
	}
}

func ExampleSubscription_Exists() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	sub := client.Subscription("subName")
	ok, err := sub.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	if !ok {
		// Subscription doesn't exist.
	}
}

func ExampleSubscription_Config() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("subName")
	config, err := sub.Config(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	fmt.Println(config)
}

func ExampleSubscription_Receive() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("subName")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		// TODO: Handle message.
		// NOTE: May be called concurrently; synchronize access to shared memory.
		m.Ack()
	})
	if err != context.Canceled {
		// TODO: Handle error.
	}
}

func ExampleSubscription_Receive_options() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("subName")
	// This program is expected to process and acknowledge messages
	// in 5 seconds. If not, Pub/Sub API will assume the message is not
	// acknowledged.
	sub.ReceiveSettings.MaxExtension = 5 * time.Second
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		// TODO: Handle message.
		m.Ack()
	})
	if err != context.Canceled {
		// TODO: Handle error.
	}
}

func ExampleSubscription_ModifyPushConfig() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("subName")
	if err := sub.ModifyPushConfig(ctx, &pubsub.PushConfig{Endpoint: "https://example.com/push"}); err != nil {
		// TODO: Handle error.
	}
}
