package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Nekodigi/pinyin-ocr-backend/config"
	"github.com/Nekodigi/pinyin-ocr-backend/handler"
	"github.com/Nekodigi/pinyin-ocr-backend/handler/util"
	"github.com/gin-gonic/gin"
)

func main() {

	//handler.Firestore()
	//pinyin formatter!
	if len(os.Args) == 2 && os.Args[1] == "test" {
		res, err := util.TranslateText("zh", "The Go Gopher is cute")
		if err != nil {
			fmt.Errorf("%+v", err)
		}
		conf := config.Load()

		fmt.Println(res)
		hres, err := http.NewRequest("GET", conf.ChargeApiUrl+"/subscribe?service_id=test&plan_id=basic&user_id=u0007", nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(hres, err)
	} else {
		engine := gin.Default()
		handler.Router(engine)
		engine.Run(":8080")
	}

}

// func translateText(targetLanguage, text string) (string, error) {
// 	// text := "The Go Gopher is cute"
// 	ctx := context.Background()

// 	lang, err := language.Parse(targetLanguage)
// 	if err != nil {
// 		return "", fmt.Errorf("language.Parse: %v", err)
// 	}

// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer client.Close()

// 	resp, err := client.Translate(ctx, []string{text}, lang, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("Translate: %v", err)
// 	}
// 	if len(resp) == 0 {
// 		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
// 	}
// 	return resp[0].Text, nil
// }
