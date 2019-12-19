package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBroadcastHandler_Happy(t *testing.T) {
	hexTxBlob := []byte("982A3CFFED5A3E37E0465D7EB97B0DD3360B413E3AC387D1197D7D961AB147F7")
	method := "POST"
	path := "/tx"
	body := bytes.NewReader(hexTxBlob)
	req := httptest.NewRequest(method, path, body)
//	req.Header.Set("x-real-ip", "0.0.0.0")

	rr := httptest.NewRecorder()
	server := &Server{
		Version: "1.0",
		Url:"tcp://data-seed-pre-0-s3.binance.org:80",
	}

	handler := http.HandlerFunc(server.TxHashHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


}