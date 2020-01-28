package main

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// SlackField ...
type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Action ...
type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	URL   string `json:"url"`
	Style string `json:"style"`
}

// Attachment ...
type Attachment struct {
	Fallback     *string       `json:"fallback"`
	Color        *string       `json:"color"`
	PreText      *string       `json:"pretext"`
	AuthorName   *string       `json:"author_name"`
	AuthorLink   *string       `json:"author_link"`
	AuthorIcon   *string       `json:"author_icon"`
	Title        *string       `json:"title"`
	TitleLink    *string       `json:"title_link"`
	Text         *string       `json:"text"`
	ImageURL     *string       `json:"image_url"`
	Fields       []*SlackField `json:"fields"`
	Footer       *string       `json:"footer"`
	FooterIcon   *string       `json:"footer_icon"`
	Timestamp    *int64        `json:"ts"`
	MarkdownIn   *[]string     `json:"mrkdwn_in"`
	Actions      []*Action     `json:"actions"`
	CallbackID   *string       `json:"callback_id"`
	ThumbnailURL *string       `json:"thumb_url"`
}

// SlackPayload ...
type SlackPayload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

// Slack ...
type Slack struct {
	WebHookURL string `json:"url"`
}

func (s *Slack) send(data *SlackPayload) []error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return []error{err}
	}
	request := gorequest.New()
	resp, _, errs := request.Post(s.WebHookURL).Send(string(jsonData)).End()
	if errs != nil {
		return errs
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.StatusCode)}
	}
	return nil
}
