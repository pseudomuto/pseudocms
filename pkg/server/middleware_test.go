package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-logr/logr/testr"
	. "github.com/pseudomuto/pseudocms/pkg/server"
)

func TestWithHTTPLogger(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	WithHTTPLogger(testr.New(t), h).ServeHTTP(w, r)
}
