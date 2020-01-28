package main

import (
	"os"
	"testing"
)

func TestZulip(t *testing.T) {

	z := Zulip{
		WebHookURL: os.Getenv("ZULIP_TEST_URL"),  // url
		Bot:        os.Getenv("ZULIP_BOT_EMAIL"), // email
		APIKey:     os.Getenv("ZULIP_API_KEY"),   // apikey
	}
	if z.WebHookURL == "" || z.Bot == "" || z.APIKey == "" {
		t.Fatal("invalid zulip config")
	}

	if err := z.send(&ZulipPayload{
		Type:    "stream",
		To:      "[프로젝트D]+서버팀",
		Subject: "unreal news",
		Content: "test message",
	}); err != nil {
		t.Error(err)
	}

	t.Log("finish zulip test")
}
