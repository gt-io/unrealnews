package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// ZulipPayload zulip payload
type ZulipPayload struct {
	Type    string `json:"type"`
	To      string `json:"to"`
	Subject string `json:"subject,omitempty"`
	Content string `json:"content"`
}

func (z *ZulipPayload) makeContent() string {
	return fmt.Sprintf(`type=%s&to=%s&subject=%s&content=%s`, z.Type, z.To, z.Subject, z.Content)
}

// Zulip ...
type Zulip struct {
	WebHookURL string `json:"url"`
	Bot        string `json:"bot"`
	APIKey     string `json:"apikey"`
}

func (z *Zulip) send(d *ZulipPayload) []error {
	request := gorequest.New()
	request.SetBasicAuth(z.Bot, z.APIKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	content := d.makeContent()
	resp, _, errs := request.Post(z.WebHookURL).Send(content).End()
	if errs != nil {
		return errs
	}

	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.StatusCode)}
	}
	return nil
}
