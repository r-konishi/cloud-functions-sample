package hellow

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHTTP(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: `{"name": ""}`, want: "Hello, World!"},
		{body: `{"name": "Gopher"}`, want: "Hello, Gopher!"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		HelloHTTP(rr, req)

		if got := rr.Body.String(); got != test.want {
			t.Errorf("HelloHTTP(%q) = %q, want %q", test.body, got, test.want)
		}
	}
}
