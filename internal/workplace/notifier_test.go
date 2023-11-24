package workplace

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	var payload struct {
		Message    string `json:"message"`
		Formatting string `json:"formatting"`
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, r.Header.Get("Authorization"), "Bearer 123")
		require.Equal(t, r.Header.Get("Content-Type"), "application/json")
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		require.Equal(t, "/group/feed", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	n := Notifier{URL: srv.URL, Token: "123"}
	err := n.Post(context.Background(), "group", "hello")
	require.NoError(t, err)
	require.Equal(t, "hello", payload.Message)
	require.Equal(t, "MARKDOWN", payload.Formatting)
}
