package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XATAB1CH/rest-api/internal/store/teststore"
	"github.com/go-playground/assert/v2"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)

}
