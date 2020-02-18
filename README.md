# Cat Slack

This small application is used to post random pictures from the [cat API](https://thecatapi.com/) to a slack channel on a regular basis.

## Requirements
1. A working API Key for the [cat API](https://thecatapi.com/)
2. A Slack hook URL. This should be formatted like `https://hooks.slack.com/services/XXXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX`

## Usage
All you need to do to run this bot is to make sure that it is running on a computer somewhere.

```
./cat-slack \
    -cat-api-key XXXXXX \
    -slack-url https://hooks.slack.com/services/XXXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX
```
By default this will start posting cat pictures to the #random channel at 9 AM every weekday.
- To change the channel that the bot posts to, use the `-channel` argument.
- To change the posting schedule, use the `-cron` argument. If you don't know how cron strings work already, check out [crontab.guru](https://crontab.guru/)