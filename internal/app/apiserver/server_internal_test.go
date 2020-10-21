package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/TerexNik/hand-meal/internal/app/model"
	"github.com/TerexNik/hand-meal/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "example@mail.ru",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		}, {
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		}, {
			name: "invalid email",
			payload: map[string]string{
				"email":    "examplemail.ru",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		}, {
			name: "invalid password",
			payload: map[string]string{
				"email":    "example@mail.ru",
				"password": "pass",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		}, {
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "wrong",
			},
			expectedCode: http.StatusUnauthorized,
		}, {
			name: "invalid email",
			payload: map[string]string{
				"email":    "wrong",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
