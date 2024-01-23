package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDataHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/{location}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
