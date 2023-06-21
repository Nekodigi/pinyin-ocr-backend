package operation

import (
	"fmt"
	"net/http"

	"github.com/Nekodigi/pinyin-ocr-backend/infrastructure/charge"
	"github.com/gin-gonic/gin"
)

type (
	Operation struct {
		Chrg *charge.Charge
	}

	SubscribeReq struct {
		UserId string
	}
)

func (u *Operation) Handle(e *gin.Engine) {
	e.POST("/subscribe", func(c *gin.Context) {
		var subscribeReq SubscribeReq
		c.BindJSON(&subscribeReq)
		url := u.Chrg.Subscribe(subscribeReq.UserId)
		if url == "" {
			fmt.Errorf("%+v", "Subscribe failed")
		} else {
			c.JSON(http.StatusOK, url)
		}
	})
}
