package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBarHandler(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(http.HandlerFunc(barHandler))
	defer s.Close()

	res, err := http.Get(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("StatusCode, Expected=%d, Recieved=%d\n", http.StatusOK, res.StatusCode)
	}
}
