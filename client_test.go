package igboapi_test

import (
	"context"
	"testing"

	"github.com/bontusss/igboapi-go"
	"github.com/stretchr/testify/assert"
)

const valid_token = ""

// func TestInvalidToken(t *testing.T) {
// 	client := igboapi.NewClient("valid_token")

// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("X_API_KEY")
// 		if token != valid_token {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		}
// 	}))

// 	defer ts.Close()

// 	client.BaseUrl = ts.URL
// 	err := client.Get(context.Background(), "/words", nil, nil)
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "request failed with status code: 401")
// }

func TestFind(t *testing.T) {
	client := igboapi.NewClient(valid_token)
	_, err := client.TermService.Find(context.Background(), "5f90c35e49f7e863e92b7264")
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	client := igboapi.NewClient(valid_token)
	_, err := client.TermService.Search(context.Background(), &igboapi.TermParams{Keyword: "bia"})
	assert.NoError(t, err)
}
