package handler

import (
	"context"
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/Nekodigi/pinyin-ocr-backend/config"
	"github.com/Nekodigi/pinyin-ocr-backend/handler/util"
	"github.com/Nekodigi/pinyin-ocr-backend/infrastructure/charge"
	"github.com/gin-gonic/gin"
)

var (
	chrg *charge.Charge
)

func init() {
	conf := config.Load()

	_ = context.Background()
	_ = vision.ImageAnnotatorClient{}
	_ = os.Open
	chrg = charge.InitCharge(conf)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Router(e *gin.Engine) {
	e.Use(CORSMiddleware())
	e.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	(&util.Util{Chrg: chrg}).Handle(e)
	// (&operation.Operation{Chrg: chrg}).Handle(e)
}
