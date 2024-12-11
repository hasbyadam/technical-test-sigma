package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hasbyadam/technical-test-sigma/entity"
	"github.com/hasbyadam/technical-test-sigma/schema/request"
	"github.com/hasbyadam/technical-test-sigma/schema/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (u *Methods) GetConfig() *entity.Config {
	return u.Config
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func generateToken(claims *entity.Claims, secretKey string) (res string, err error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	res, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return
	}
	return
}

func (u *Methods) Register(ctx context.Context, req request.Register) (res response.Register, err error) {
	hashPassword, err := getHashPassword(req.Password)
	if err != nil {
		return
	}

	userId := uuid.New()
	if err = u.Stores.Register(ctx, entity.UserDetails{
		Id:         userId,
		Email:      req.Email,
		Password:   hashPassword,
		FullName:   req.FullName,
		LegalName:  req.LegalName,
		NIK:        req.NIK,
		BirthPlace: req.BirthPlace,
		BirthDate:  req.BirthDate,
		Salary:     req.Salary,
		KTP:        req.KTP,
		KTPSelfie:  req.KTPSelfie,
		CreatedAt:  time.Now().Unix(),
		CreatedBy:  userId,
		IsVerifed:  false,
	}); err != nil {
		zap.S().Info(err)
		return
	}
	claims := &entity.Claims{
		UserId: userId,
		Admin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	res.Token, err = generateToken(claims, u.Config.Auth.JwtSecretKey)
	if err != nil {
		zap.S().Info(err)
	}
	return
}

func (u *Methods) Login(ctx context.Context, req request.Login) (res response.Login, err error) {

	userDetails, err := u.Stores.GetUserDetails(ctx, req.Email)
	if err != nil {
		return
	}
	if !checkPassword(userDetails.Password, req.Password) {
		return res, errors.New("invalid password")
	}

	claims := &entity.Claims{
		UserId: userDetails.Id,
		Admin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	res.Token, err = generateToken(claims, u.Config.Auth.JwtSecretKey)
	if err != nil {
		zap.S().Info(err)
	}
	return
}

func (u *Methods) ParseTokenToClaims(tokenString string) (err error) {
	if _, err = jwt.ParseWithClaims(tokenString, &u.Claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(u.Config.Auth.JwtSecretKey), nil
		},
	); err != nil {
		zap.S().Info(err)
		return
	}
	return
}
