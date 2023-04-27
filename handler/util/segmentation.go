package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ego/gse"
)

type (
	Segmentation struct {
	}

	SegmentationReq struct {
		Text string
	}
)

var testDict string
var seg gse.Segmenter

func init() {
	seg, _ = gse.NewEmbed("zh, word 20 n"+testDict, "en")
}

func (u *Segmentation) Handle(e *gin.Engine) {

	e.POST("/segmentation", func(c *gin.Context) {
		var segmentationReq SegmentationReq
		c.BindJSON(&segmentationReq)

		target, _ := TranslateText("zh", segmentationReq.Text)
		fmt.Println(target)
		s1 := seg.Cut(target, true)

		// fmt.Println("stop: ", seg.Stop(s1))
		c.JSON(http.StatusAccepted, seg.Stop(s1))
	})

}
