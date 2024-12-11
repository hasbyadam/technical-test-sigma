package request

import "gopkg.in/guregu/null.v4"

type Register struct {
	Email      null.String `json:"email"`
	Password   string      `json:"password"`
	FullName   null.String `json:"fullName"`
	LegalName  null.String `json:"legalName"`
	NIK        null.String `json:"nik"`
	BirthPlace null.String `json:"birthPlace"`
	BirthDate  null.String `json:"birthDate"`
	Salary     null.String `json:"salary"`
	KTP        null.String `json:"ktp"`
	KTPSelfie  null.String `json:"ktpSelfie"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
