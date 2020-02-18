package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/bobcob7/cat-slack/pkg/catbot"
	"gopkg.in/robfig/cron.v2"
)

const (
	catAPIKeyKey     = "cat-api-key"
	catAPIKeyUsage   = "API key to access cat API"
	catAPIKeyDefault = ""

	slackURLKey     = "slack-url"
	slackURLUsage   = "Slack webhook URL to send the images to"
	slackURLDefault = ""

	cronStringKey     = "cron"
	cronStringUsage   = "Cron string for how often to trigger the cat API"
	cronStringDefault = "0 9 * * 1-5"

	channelKey     = "channel"
	channelUsage   = "Slack channel to send image to"
	channelDefault = "random"

	quietKey     = "q"
	quietUsage   = "Run server quietly"
	quietDefault = false
)

var (
	catAPIKey  string
	slackURL   string
	cronString string
	channel    string
	quiet      bool
)

func init() {
	flag.StringVar(&catAPIKey, catAPIKeyKey, catAPIKeyDefault, catAPIKeyUsage)
	flag.StringVar(&slackURL, slackURLKey, slackURLDefault, slackURLUsage)
	flag.StringVar(&cronString, cronStringKey, cronStringDefault, cronStringUsage)
	flag.StringVar(&channel, channelKey, channelDefault, channelUsage)
	flag.BoolVar(&quiet, quietKey, quietDefault, quietUsage)
}

func validateArgs() error {
	if len(catAPIKey) <= 0 {
		return errors.New("Missing cat API key")
	}
	if len(slackURL) <= 0 {
		return errors.New("Missing slack URL")
	}
	if len(cronString) <= 0 {
		return errors.New("Empty cron string")
	}
	if len(channel) <= 0 {
		return errors.New("Empty channel")
	}
	return nil
}

func main() {
	flag.Parse()
	err := validateArgs()
	if err != nil {
		fmt.Println("Error running cat API:", err)
		return
	}
	ctx := context.Background()
	cat := catbot.Cat{
		APIKey: catAPIKey,
	}
	if err := cat.Verify(); err != nil {
		fmt.Println("Failed to initialized Cat API:", err)
		return
	}
	bot := catbot.Slack{
		URL:     slackURL,
		Cat:     cat,
		Channel: channel,
	}
	if err := bot.Verify(); err != nil {
		fmt.Println("Failed to initialized Slack bot:", err)
		return
	}
	c := cron.New()
	if _, err := c.AddFunc(cronString, bot.SendRandomCatImage); err != nil {
		fmt.Println("Failed to scheduler:", err)
		return
	}
	c.Start()
	<-ctx.Done()
}
