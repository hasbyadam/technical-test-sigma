package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Id    uuid.UUID `json:"id"`
	Admin bool      `json:"admin"`
	jwt.RegisteredClaims
}

type UserDetails struct {
	Id         uuid.UUID
	Email      string
	Passowrd   string
	FullName   string
	LegalName  string
	NIK        string
	BirthPlace string
	BirthDate  string
	Salary     string
	KTP        string
	KTPSelfie  string
	CreatedAt  int64
	CreatedBy  uuid.UUID
	IsVerifed  bool
}
