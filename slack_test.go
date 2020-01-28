package main

import (
	"os"
	"testing"
)

func TestSlack(t *testing.T) {

	s := Slack{
		WebHookURL: os.Getenv("SLACK_TEST_URL"),
	}

	if s.WebHookURL == "" {
		t.Fatal("url is empty")
	}

	if err := s.send(&SlackPayload{
		Text: "This is test code for slack.",
	}); err != nil {
		for _, v := range err {
			t.Error(v)
		}
	}

	t.Log("finish slack test")
}
