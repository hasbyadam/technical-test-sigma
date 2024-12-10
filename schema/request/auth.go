package request

type Register struct {
	Email      string `json:"email"`
	Passowrd   string `json:"password"`
	FullName   string `json:"fullName"`
	LegalName  string `json:"legalName"`
	NIK        string `json:"nik"`
	BirthPlace string `json:"birthPlace"`
	BirthDate  string `json:"birthDate"`
	Salary     string `json:"salary"`
	KTP        string `json:"ktp"`
	KTPSelfie  string `json:"ktpSelfie"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
