package telegram

import (
	"strings"

	"wechatbot/config"
	"wechatbot/openai"

	log "github.com/sirupsen/logrus"
)

func Handle(msg string) *string {
	model := config.GetModel()
	requestText := strings.TrimSpace(msg)
	reply, err := openai.Completions(requestText, *model)
	if err != nil {
		log.Error(err)
	}
	return reply
}
