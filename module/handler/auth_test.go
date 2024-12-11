package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	appInit "github.com/hasbyadam/technical-test-sigma/init"
	"github.com/hasbyadam/technical-test-sigma/module/store"
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockConfig = appInit.SetupMainConfig()
)

func TestCreateUserSuccess(t *testing.T) {
	requestJson := fmt.Sprintf(`{"email": "email%d@mail.com","password": "hasby123"}`, rand.Int())

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/register")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}

func TestCreateUserBadRequest(t *testing.T) {
	requestJson := `{"email": "hasby@mail.com","password": "hasby123"}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/register")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.Error(t, handler.Register(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}

func TestUserLoginBadRequest(t *testing.T) {
	requestJson := fmt.Sprintf(`{"email": "email%d@mail.com","password": "aaaaaaa"}`, rand.Int())

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/login")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}

func TestUserLoginSuccess(t *testing.T) {
	requestJson := `{"email": "hasby@mail.com","password": "hasby123"}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/login")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.Error(t, handler.Register(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}
