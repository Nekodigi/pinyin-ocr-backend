package charge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Nekodigi/pinyin-ocr-backend/config"
	"github.com/Nekodigi/pinyin-ocr-backend/consts"
	"github.com/gin-gonic/gin"
)

var c *Charge

type (
	Charge struct {
		url       string
		serviceId string
	}

	StatusRes struct {
		Status string `json:"status"`
	}

	UrlRes struct {
		Url string `json:"url"`
	}
)

func InitCharge(conf *config.Config) *Charge {
	c = &Charge{
		url:       conf.ChargeApiUrl,
		serviceId: conf.ServiceId,
	}
	return c
}

func (c *Charge) UseQuota(ctx *gin.Context, userId string, amount float64) bool {
	//fmt.Println(c, userId, amount)
	//fmt.Println(fmt.Sprintf("%s/use_quota/%s/%s?amount=%f", c.url, c.serviceId, userId, amount))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/use_quota/%s/%s?amount=%f", c.url, c.serviceId, userId, amount), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	client := &http.Client{}
	resp_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp_.Body.Close()
	body, _ := io.ReadAll(resp_.Body)
	var resp StatusRes
	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Printf("Error unmarshal %+v\n", err)
	}
	if resp_.Status != "200 OK" {
		fmt.Printf("Error using quota: %s\n", resp_.Status)
		ctx.JSON(http.StatusBadRequest, resp)
		return false
	}
	if resp.Status != consts.OK {
		ctx.JSON(http.StatusBadRequest, resp)
		return false
	}
	return true
}

func (c *Charge) Subscribe(userId string) string {
	//fmt.Println(c, userId, amount)
	//fmt.Println(fmt.Sprintf("%s/subscribe/%s/%s/%s", c.url, c.serviceId, userId, "basic"))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/subscribe/%s/%s/%s", c.url, c.serviceId, userId, "basic"), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	client := &http.Client{}
	resp_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp_.Body.Close()
	body, _ := io.ReadAll(resp_.Body)
	var resp UrlRes
	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Printf("Error unmarshal %+v\n", err)
	}
	if resp_.Status != "200 OK" {
		fmt.Printf("Error getting subscription url: %s\n", resp_.Status)
		return ""
	}
	return resp.Url
}
