package store

import (
	"context"

	"github.com/hasbyadam/technical-test-sigma/entity"
	"go.uber.org/zap"
)

func (c *Client) Register(ctx context.Context, req entity.UserDetails) (err error) {
	qs := `INSERT INTO sigma_test.users
(id, password, email, full_name, legal_name, nik, birth_place, birth_date, salary, ktp, ktp_selfie, created_at, created_by, is_verified)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	if _, err := c.mysql.ExecContext(ctx, qs,
		req.Id,
		req.Passowrd,
		req.Email,
		req.FullName,
		req.LegalName,
		req.NIK,
		req.BirthPlace,
		req.BirthDate,
		req.Salary,
		req.KTP,
		req.KTPSelfie,
		req.CreatedAt,
		req.CreatedBy,
		req.IsVerifed,
	); err != nil {
		zap.S().Info(err)
	}

	return
}

func (c *Client) GetUserDetails(ctx context.Context, email string) (res entity.UserDetails, err error) {
	qs := `SELECT  id, password, email, full_name, legal_name, nik, birth_place, birth_date, salary, ktp, ktp_selfie, created_at, created_by, is_verified
FROM sigma_test.users where email = ?`

	row := c.mysql.QueryRowContext(ctx, qs, email)
	if err = row.Scan(
		&res.Id,
		&res.Passowrd,
		&res.Email,
		&res.FullName,
		&res.LegalName,
		&res.NIK,
		&res.BirthPlace,
		&res.BirthDate,
		&res.Salary,
		&res.KTP,
		&res.KTPSelfie,
		&res.CreatedAt,
		&res.CreatedBy,
		&res.IsVerifed,
	); err != nil {
		zap.S().Info(err)
	}

	return
}
