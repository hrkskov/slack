package main

import (
	"fmt"
	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")

	channelID, _, err := api.PostMessage(
		"CHANNEL_ID",
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionMetadata(slack.Metadata{EventType: "some_event_type", EventPayload: map[string]string{"some_event_payload": "some body"}}),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	conversationHistory, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{ChannelID: channelID})
	if err != nil {
		fmt.Printf("Got metadata: %s\n", conversationHistory.Messages[0].Metadata)
		return
	}
}
