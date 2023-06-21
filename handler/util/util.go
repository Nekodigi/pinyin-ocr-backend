package util

import (
	"github.com/Nekodigi/pinyin-ocr-backend/infrastructure/charge"
	"github.com/gin-gonic/gin"
)

type (
	Util struct {
		Chrg *charge.Charge
	}
)

func (u *Util) Handle(e *gin.Engine) {
	u.HandleOcr(e)
	u.HandleSegmentation(e)
	u.HandleTranslate(e)
}
