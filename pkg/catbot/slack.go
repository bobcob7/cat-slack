package catbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	welcomeMessage = "Cat bot reporting for duty! Nya!"
)

type Slack struct {
	URL     string
	Cat     Cat
	Channel string
}

type SlackMessage struct {
	Channel string `json:"channel`
	Text    string `json:"text"`
}

func (s Slack) Verify() error {
	msg := SlackMessage{
		Channel: s.Channel,
		Text:    welcomeMessage,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Bad status %d: %s", resp.StatusCode, resp.Status)
	}
	return nil
}

func (s Slack) SendImage(channel, imageURL string) error {
	msg := SlackMessage{
		Channel: channel,
		Text:    imageURL,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Bad status %d: %s", resp.StatusCode, resp.Status)
	}
	return nil
}

func (s Slack) SendRandomCatImage() {
	url, err := s.Cat.GetRandomURL()
	if err == nil {
		err = s.SendImage(s.Channel, url)
	}
	if err != nil {
		log.Println("Failed to send cat image", err)
	}
}
