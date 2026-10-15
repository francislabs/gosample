package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello world" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world")
	}
}

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	indexHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("got status %d expected %d", rec.Code, http.StatusOK)
	}
	if !strings.Contains(rec.Body.String(), "hello world") {
		t.Errorf("response body does not contain expected text")
	}
}
