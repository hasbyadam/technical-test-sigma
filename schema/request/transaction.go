package request

import "github.com/google/uuid"

type RequestTransaction struct {
	CreditLimitId uuid.UUID `json:"creditLimitId"`
	AssetName     string    `json:"assetName"`
	Installment   int64     `json:"installment"`
}
