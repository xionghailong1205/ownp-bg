package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"ownp-bg/internal/service"

	"github.com/gin-gonic/gin"
)

type mockHealthService struct {
	status service.HealthStatus
	err    error
}

func (m mockHealthService) Status(_ context.Context) (service.HealthStatus, error) {
	return m.status, m.err
}

func TestGetStatusSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := NewHealthHandler(mockHealthService{
		status: service.HealthStatus{
			Status:       "ok",
			DatabaseTime: time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
		},
	})

	r := gin.New()
	r.GET("/api/v1/health", h.GetStatus)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var payload map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &payload); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected status ok, got %v", payload["status"])
	}
}

func TestGetStatusDatabaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := NewHealthHandler(mockHealthService{
		err: errors.New("db failed"),
	})

	r := gin.New()
	r.GET("/api/v1/health", h.GetStatus)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
