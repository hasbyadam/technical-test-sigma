package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hasbyadam/technical-test-sigma/entity"
	"github.com/hasbyadam/technical-test-sigma/schema/request"
	"github.com/hasbyadam/technical-test-sigma/schema/response"
	"go.uber.org/zap"
)

func (u *Methods) RequestTransaction(ctx context.Context, req request.RequestTransaction) (res response.RequestTransaction, err error) {
	
	limit, err := u.Stores.GetUserLimitBalance(ctx, u.Claims.UserId, req.CreditLimitId)
	if err != nil {
		zap.S().Info(err)
		return
	}

	if limit.LimitBalance < req.Installment {
		return res, errors.New("insufficient balance")
	}

	res.ContractNumber = fmt.Sprintf("CN-%d", time.Now().UnixMilli())
	if err = u.Stores.InsertTransactions(ctx, entity.TransactionDetails{
		Id:             uuid.New(),
		ContractNumber: res.ContractNumber,
		Installment:    req.Installment,
		AssetName:      req.AssetName,
		CreatedAt:      time.Now().Unix(),
		OTR:            1,
		Interest:       1,
		AdminFee:       1,
		CreatedBy:      u.Claims.UserId.String(),
		IsSettled:      false,
		CreditLimitId:  req.CreditLimitId,
	}); err != nil {
		return
	}

	return
}
