package handler

import (
	"bytes"
	"genxoft.dev/internal/store/sqlitestore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenHandler_ServeHTTP(t *testing.T) {
	jsonBody := []byte(`{
		"token": "test_token",
		"timezone": "Asia/Tbilisi",
		"ip": "192.168.1.1"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "/api/fcm-token", bodyReader)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	s, teardown := sqlitestore.TestStore(t)
	defer teardown("fcm_settings")

	handler := FcmTokenHandler(s.FcmToken())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fc, err := s.FcmToken().Find("test_token")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "test_token", fc.Token)
}
