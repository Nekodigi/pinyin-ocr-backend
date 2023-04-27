package main

import (
	"os"

	"github.com/Nekodigi/pinyin-ocr-backend/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	//handler.Firestore()
	if len(os.Args) == 2 && os.Args[1] == "test" {
		// err := util.DetectText("sample.png")
		// if err != nil {
		// 	fmt.Errorf("%+v", err)
		// }
	} else {
		engine := gin.Default()
		handler.Router(engine)
		engine.Run(":8080")
	}

}
