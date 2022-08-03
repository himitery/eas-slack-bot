package expo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/himitery/eas-slack-bot/internal/domain"
	"github.com/himitery/eas-slack-bot/internal/service/secret"
	"github.com/himitery/eas-slack-bot/internal/service/slack"
	"net/http"
)

func SubmitHook(context *gin.Context) {
	var json domain.SubmitResult
	if err := context.ShouldBindJSON(&json); err != nil {
		fmt.Println("Failed to bind json: " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if signature := context.Request.Header.Get("expo-signature"); len(signature) == 0 || signature != secret.GetHashCode(json) {
		fmt.Println("Signatures didn't match")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Signatures didn't match"})
		return
	}

	slack.SendMessage(slack.SendMessageRequest{
		Type:     "submit",
		Platform: json.Platform,
		Status:   json.Status,
	})

	context.JSON(http.StatusOK, gin.H{})
}
