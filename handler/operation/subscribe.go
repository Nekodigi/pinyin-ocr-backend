package operation

import (
	"github.com/Nekodigi/pinyin-ocr-backend/infrastructure/charge"
)

type (
	Operation struct {
		Chrg *charge.Charge
	}

	SubscribeReq struct {
		UserId string
	}
)
