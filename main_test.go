package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
	if !strings.Contains(string(body), "Hello, CI/CD!") {
		t.Errorf("Expected body to contain 'Hello, CI/CD!', got %s", string(body))
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
	if string(body) != "OK" {
		t.Errorf("Expected body 'OK', got %s", string(body))
	}
}

func TestAboutHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()

	aboutHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
	expected := "This is a basic Go server with multiple endpoints."
	if string(body) != expected {
		t.Errorf("Expected body '%s', got %s", expected, string(body))
	}
}
