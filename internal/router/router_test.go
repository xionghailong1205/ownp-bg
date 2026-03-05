package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"ownp-bg/internal/handler"
	"ownp-bg/internal/service"
)

type mockHealthService struct{}

func (m mockHealthService) Status(_ context.Context) (service.HealthStatus, error) {
	return service.HealthStatus{
		Status:       "ok",
		DatabaseTime: time.Now(),
	}, nil
}

func TestOpenAPIRoute(t *testing.T) {
	h := handler.NewHealthHandler(mockHealthService{})
	r := New(h)

	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type: %s", got)
	}
}
