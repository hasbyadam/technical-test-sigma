package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/hasbyadam/technical-test-sigma/entity"
	"go.uber.org/zap"
)

func (c *Client) InsertTransactions(ctx context.Context, req entity.TransactionDetails) (err error) {
	LockTables(c.mysql, "sigma_test.transactions")
	defer UnlockTables(c.mysql)
	qs := `INSERT INTO sigma_test.transactions
(id, contract_number, otr, admin_fee, installment, interest, asset_name, created_by, created_at, is_settled, credit_limit_id)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	if _, err = c.mysql.ExecContext(ctx, qs,
		req.Id,
		req.ContractNumber,
		req.OTR,
		req.AdminFee,
		req.Installment,
		req.Interest,
		req.AssetName,
		req.CreatedBy,
		req.CreatedAt,
		req.IsSettled,
		req.CreditLimitId,
	); err != nil {
		zap.S().Info(err)
		return
	}

	return
}

func (c *Client) GetUserLimitBalance(ctx context.Context, userId, creditLimitId uuid.UUID) (res entity.UserLimitBalance, err error) {
	qs := `SELECT user_id, credit_limit_id, tenor, limit_balance
FROM sigma_test.user_limit_balances where user_id = ? and credit_limit_id = ?`
	row := c.mysql.QueryRowContext(ctx, qs, userId, creditLimitId)
	
	if err = row.Scan(
		&res.UserId,
		&res.CreditLimitId,
		&res.Tenor,
		&res.LimitBalance,
	); err != nil {
		return
	}

	return
}
