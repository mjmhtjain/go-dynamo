package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/health"), nil)
	res := httptest.NewRecorder()

	HealthHandler(res, req)

	actualStatusCode := res.Result().StatusCode
	expectedStatusCode := http.StatusOK
	if actualStatusCode != expectedStatusCode {
		t.Errorf("exp: %v, actual: %v", expectedStatusCode, actualStatusCode)
	}
}
