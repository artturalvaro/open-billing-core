package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSubscriptionAPI(t *testing.T) {
	body := map[string]string{
		"customerID": "cust_test",
		"planID":     "plan_basic",
	}

	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/subscriptions", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := setupServer()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
}
