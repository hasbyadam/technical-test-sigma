package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Claims struct {
	UserId uuid.UUID `json:"userId"`
	Admin  bool      `json:"admin"`
	jwt.RegisteredClaims
}

type UserDetails struct {
	Id         uuid.UUID
	Email      null.String
	Password   string
	FullName   null.String
	LegalName  null.String
	NIK        null.String
	BirthPlace null.String
	BirthDate  null.String
	Salary     null.String
	KTP        null.String
	KTPSelfie  null.String
	CreatedAt  int64
	CreatedBy  uuid.UUID
	IsVerifed  bool
}
