package entity

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type TransactionDetails struct {
	Id             uuid.UUID
	ContractNumber string
	OTR            int64
	AdminFee       int64
	Installment    int64
	Interest       int64
	AssetName      string
	CreatedBy      string
	CreatedAt      int64
	CreditLimitId  uuid.UUID
	IsSettled      bool
}

type UserCreditLimit struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Tenor     string
	Credit    int64
	CreatedAt int64
	UpdatedAt null.Int
}

type UserLimitBalance struct {
	UserId        uuid.UUID
	CreditLimitId uuid.UUID
	Tenor         string
	LimitBalance  int64
}
