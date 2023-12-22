package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type Config struct {
	SLACK_BOT_TOKEN string
	CHANNEL_ID      string
}

func setUpSlackChannelConfig(slackBotToken, ChannelID string) Config {
	return Config{
		SLACK_BOT_TOKEN: slackBotToken,
		CHANNEL_ID:      ChannelID,
	}
}
func checkFileName() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
}

func createSlackClient(c *Config) *slack.Client {
	return slack.New(c.SLACK_BOT_TOKEN)
}

func isValidFile() bool {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return false
	}

	return true

}

func uploadFile(fileName string, c *Config, api *slack.Client) {

	params := slack.FileUploadParameters{
		Channels: []string{c.CHANNEL_ID},
		File:     fileName,
	}

	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)

}
func main() {

	if !isValidFile() {
		return
	}

	config := setUpSlackChannelConfig(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("CHANNEL_ID"))

	api := createSlackClient(&config)

	checkFileName()

	uploadFile(os.Args[1], &config, api)
}
