package healthhandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func TestHello(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HelloWorldHandler))
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", res.StatusCode)
	}

	var response HelloResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	require.Equal(t, "Hello World", response.Message)
	require.NotEmpty(t, response.Message)
}
