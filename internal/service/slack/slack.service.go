package slack

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"os"
)

var client *slack.Client

type SendMessageRequest struct {
	Type           string
	Platform       string
	Status         string
	BuildUrl       *string // required if type is build
	AppVersion     *string // required if type is build
	SdkVersion     *string // required if type is build
	BuildProfile   *string // required if type is build
	ReleaseChannel *string // required if type is build
}

func getSlackClient() *slack.Client {
	if client == nil {
		client = slack.New(os.Getenv("SLACK_TOKEN"))
	}
	return client
}

func passNilValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func SendMessage(message SendMessageRequest) {
	title := func(messageType string) string {
		status := func(status string) string {
			if status == "errored" {
				return "실패하였습니다."
			} else if status == "finished" {
				return "완료되었습니다."
			} else {
				return "취소되었습니다."
			}
		}(message.Status)

		if messageType == "build" {
			return fmt.Sprintf("*[%s](%s) 빌드가 %s*\n", message.Platform, passNilValue(message.AppVersion), status)
		} else {
			return fmt.Sprintf("[%s] 배포가 %s\n", message.Platform, status)
		}
	}(message.Type)

	info := func(messageType string, status string) string {
		if messageType == "build" {
			infoMessage := fmt.Sprintf("SDK Version: %s\nBuild Profile: %s\n", passNilValue(message.SdkVersion), passNilValue(message.BuildProfile))
			if status == "build" {
				infoMessage = infoMessage + fmt.Sprintf("Release Channel: %s\nBuild Url: %s\n", passNilValue(message.ReleaseChannel), passNilValue(message.BuildUrl))
			}
			return infoMessage
		} else {
			return ""
		}
	}(message.Type, message.Status)

	attachment := slack.Attachment{
		Pretext:    title,
		Text:       info,
		MarkdownIn: []string{"pretext", "info"},
	}

	if _, _, err := getSlackClient().PostMessage(
		os.Getenv("SLACK_CHANNEL_NAME"),
		slack.MsgOptionAttachments(attachment),
	); err != nil {
		log.Println("Failed to post slack message: ", err.Error())
		return
	}
}
