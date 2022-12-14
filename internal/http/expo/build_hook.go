package expo

import (
	"github.com/gin-gonic/gin"
	"github.com/himitery/eas-slack-bot/internal/domain"
	"github.com/himitery/eas-slack-bot/internal/service/secret"
	"github.com/himitery/eas-slack-bot/internal/service/slack"
	"log"
	"net/http"
)

func BuildHook(context *gin.Context) {
	var json domain.BuildResult
	if err := context.ShouldBindJSON(&json); err != nil {
		log.Println("Failed to bind json: " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if signature := context.Request.Header.Get("expo-signature"); len(signature) == 0 || signature != secret.GetHashCode(json) {
		log.Println("Signatures didn't match")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Signatures didn't match"})
		return
	}

	slack.SendMessage(slack.SendMessageRequest{
		Type:           "build",
		Platform:       json.Platform,
		Status:         json.Status,
		BuildUrl:       json.Artifacts.BuildUrl,
		AppVersion:     &json.Metadata.AppVersion,
		SdkVersion:     &json.Metadata.SdkVersion,
		BuildProfile:   &json.Metadata.BuildProfile,
		ReleaseChannel: json.Metadata.ReleaseChannel,
	})

	context.JSON(http.StatusOK, gin.H{})
}
