package main

import (
	"context"

	"github.com/rotemtam/action-workplace-notify/internal/workplace"
	"github.com/sethvargo/go-githubactions"
)

const DefaultBaseURL = "https://graph.workplace.com"

func main() {
	act := githubactions.New()
	act.Infof("Hello, world!")
	n := workplace.Notifier{
		URL:   DefaultBaseURL,
		Token: act.GetInput("access-token"),
	}
	if n.Token == "" {
		act.Fatalf("missing token")
	}
	var (
		groupID = act.GetInput("group-id")
		message = act.GetInput("message")
	)
	if groupID == "" {
		act.Fatalf("missing group-id")
	}
	if message == "" {
		act.Fatalf("missing message")
	}
	if err := n.Post(context.Background(), groupID, message); err != nil {
		act.Fatalf("failed to post message: %v", err)
	}
	act.Infof("message posted")
}
