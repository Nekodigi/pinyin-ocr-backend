package util

import (
	"context"
	"fmt"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
)

type (
	OCR struct {
	}
)

func (u *OCR) Handle(e *gin.Engine) {
	e.POST("/ocr", func(c *gin.Context) {
		// Using 'ShouldBind'
		formFile, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		ctx := context.Background()
		client, err := vision.NewImageAnnotatorClient(ctx)
		if err != nil {
			fmt.Errorf("%+v", err)
		}

		f, err := formFile.Open()
		if err != nil {
			fmt.Errorf("%+v", err)
		}
		defer f.Close()

		image, err := vision.NewImageFromReader(f)
		if err != nil {
			fmt.Errorf("%+v", err)
		}
		//fmt.Println("form error", formFile)
		annotations, err := client.DetectTexts(ctx, image, nil, 10)
		if err != nil {
			fmt.Errorf("%+v", err)
		}

		if len(annotations) == 0 {
			fmt.Println("Text not found")
			c.JSON(http.StatusNoContent, "Text not found")

		} else {
			fmt.Println("Text:")
			fmt.Printf("%q\n", annotations[0].Description)
			c.JSON(http.StatusAccepted, annotations[0].Description)
		}
	})
}
