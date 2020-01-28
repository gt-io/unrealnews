package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Bot ...
type Bot struct {
	zulip *Zulip
	slack *Slack
}

func initBot(path string) (*Bot, error) {
	// load config from file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf struct {
		ZulipConfig *struct {
			WebHookURL string `json:"webhookurl"`
			BotEmail   string `json:"bot_email"`
			APIKey     string `json:"api_key"`
		} `json:"zulip"`
		SlackConfig *struct {
			WebHookURL string `json:"webhookurl"`
		} `json:"slack"`
	}
	if err := json.Unmarshal(bytes, &conf); err != nil {
		return nil, err
	}
	log.Print(conf)

	//
	var bot Bot
	if conf.ZulipConfig != nil && conf.ZulipConfig.WebHookURL != "" {
		bot.zulip = &Zulip{
			WebHookURL: conf.ZulipConfig.WebHookURL,
			Bot:        conf.ZulipConfig.BotEmail,
			APIKey:     conf.ZulipConfig.APIKey,
		}
	}
	if conf.SlackConfig != nil && conf.SlackConfig.WebHookURL != "" {
		bot.slack = &Slack{
			WebHookURL: conf.SlackConfig.WebHookURL,
		}
	}

	return &bot, nil
}

func (b *Bot) sendBot(data string) error {
	log.Println("send bot message..", data)
	if b.zulip != nil {
		// log.Println("send bot zulip message..", b.zulip)

		if err := b.zulip.send(&ZulipPayload{
			Type:    "stream",
			To:      "[프로젝트D]+서버팀",
			Subject: "unreal news",
			Content: data,
		}); err != nil {
			log.Println(err)
			return err[0]
		}

	}
	if b.slack != nil {
		// log.Println("send bot slack message..", b.slack)

		if err := b.slack.send(&SlackPayload{
			Text: data,
		}); err != nil {
			log.Println(err)
			return err[0]
		}

	}
	return nil
}
