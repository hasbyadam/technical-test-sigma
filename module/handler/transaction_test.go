package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hasbyadam/technical-test-sigma/module/store"
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRequestTransactionSuccess(t *testing.T) {
	requestJson := `{"creditLimitId": "82e47570-fcd4-4115-9e80-a9b13c9bb124","assetName": "Test","installment": 10000}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transaction")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}

func TestRequestTransactionBadRequest(t *testing.T) {
	requestJson := `{"creditLimitId": "","assetName": "Test","installment": 10000}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transaction")
	handler := &Handler{
		Usecase: usecase.New(store.New(mockConfig), mockConfig),
	}

	if assert.Error(t, handler.Register(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, requestJson, rec.Body.String())
	}
}
