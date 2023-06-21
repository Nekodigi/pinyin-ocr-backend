package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ego/gse"
)

type (
	SegmentationReq struct {
		UserId string
		Text   string
	}
)

var testDict string
var seg gse.Segmenter

func init() {
	seg, _ = gse.NewEmbed("zh, word 20 n"+testDict, "en")
}

func (u *Util) HandleSegmentation(e *gin.Engine) {

	e.POST("/segmentation", func(c *gin.Context) {
		var segmentationReq SegmentationReq
		c.BindJSON(&segmentationReq)
		fmt.Println(segmentationReq.Text)
		//0.004 per char
		if !u.Chrg.UseQuota(c, segmentationReq.UserId, float64(len(segmentationReq.Text))*0.004) {
			return
		}

		target, err := TranslateText("zh", segmentationReq.Text)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(target)
		s1 := seg.Cut(target, true)

		fmt.Println("stop: ", seg.Stop(s1))
		c.JSON(http.StatusAccepted, seg.Stop(s1))
	})

}
