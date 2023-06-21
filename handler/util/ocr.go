package util

import (
	"context"
	"fmt"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
)

type ()

func (u *Util) HandleOcr(e *gin.Engine) {
	e.POST("/ocr", func(c *gin.Context) {
		// Using 'ShouldBind'
		formFile, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		user_id := c.PostForm("user_id")

		//charge 0.3
		if !u.Chrg.UseQuota(c, user_id, 0.3) {
			return
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
			res := []string{}
			for _, annotation := range annotations[1:] {
				res = append(res, annotation.Description)
			}
			c.JSON(http.StatusAccepted, res)
		}
	})
}

// func DetectText(file string) error {
// 	ctx := context.Background()

// 	client, err := vision.NewImageAnnotatorClient(ctx)
// 	if err != nil {

// 		return err
// 	}
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	image, err := vision.NewImageFromReader(f)
// 	if err != nil {
// 		return err
// 	}

// 	annotations, err := client.DetectTexts(ctx, image, nil, 10)
// 	if err != nil {
// 		return err
// 	}

// 	if len(annotations) == 0 {
// 		fmt.Println("No text found.")
// 	} else {
// 		fmt.Println("Text:")
// 		for _, annotation := range annotations {
// 			fmt.Printf("%q\n", annotation.Description)
// 		}
// 	}

// 	return nil
// }
