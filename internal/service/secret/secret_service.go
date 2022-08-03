package secret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/himitery/eas-slack-bot/internal/domain"
	"os"
)

func GetHashCode[T domain.BuildResult | domain.SubmitResult](body T) string {
	hash := hmac.New(sha1.New, []byte(os.Getenv("SECRET_WEBHOOK_KEY")))
	jsonBytes, _ := json.Marshal(body)
	hash.Write(jsonBytes)
	return "sha1=" + hex.EncodeToString(hash.Sum(nil))
}
