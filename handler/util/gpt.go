package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type (
	GPT struct {
		OpenAI *openai.Client
	}
	PinyinReq struct {
		Text string
	}
)

func (u *GPT) Handle(e *gin.Engine) {
	e.POST("/gpt", func(c *gin.Context) {
		var pinyinReq PinyinReq
		c.BindJSON(&pinyinReq)
		fmt.Printf("Pinyin: %s", pinyinReq.Text)

		messages := []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "generate html code for pinyin as ruby:"},
			{Role: openai.ChatMessageRoleUser, Content: "generate html code for pinyin as ruby:很高兴你觉得有趣！有什么可以帮助你的吗？"},
			{Role: openai.ChatMessageRoleAssistant, Content: "<span lang='zh-CN'><ruby>很高兴<rt>hěn gāo xìng</rt></ruby> <ruby>你<rt>nǐ</rt></ruby> <ruby>觉得<rt>jué de</rt></ruby> <ruby>有趣<rt>yǒu qù</rt></ruby>！<ruby>有什么<rt>yǒu shén me</rt></ruby> <ruby>可以<rt>kě yǐ</rt></ruby> <ruby>帮助<rt>bāng zhù</rt></ruby> <ruby>你<rt>nǐ</rt></ruby> <ruby>的<rt>de</rt></ruby> <ruby>吗<rt>ma</rt></ruby>？</span>"},
			{Role: openai.ChatMessageRoleUser, Content: fmt.Sprintf("generate html code for pinyin as ruby:%s", pinyinReq.Text)},
		}

		resp, err := u.OpenAI.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			c.JSON(http.StatusNoContent, "Text not found")
		}
		fmt.Println(resp)
		c.JSON(http.StatusAccepted, resp.Choices[0].Message.Content)
	})
}
