package igboapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

const baseUrl = "https://igboapi.com/api/v1/"

type Client struct {
	TermService TermService

	HTTPClient *http.Client
	Token      string
	BaseUrl    string
}

func NewClient(token string) *Client {
	var client Client
	client.TermService = &termService{client: &client}
	client.BaseUrl = baseUrl
	client.Token = token

	return &client
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}

	return http.DefaultClient
}

// func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
// 	req = req.WithContext(ctx)
// 	req.Header.Set("X_API_KEY", c.Token)
// 	req.Header.Set("Content-Type", "application/json")
// 	return c.httpClient().Do(req)
// }

func (c *Client) get(ctx context.Context, path string, params url.Values, v interface{}) error {
	fmt.Println("ENTERING GET")
	reqURL := c.BaseUrl + path
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP request: %w", err)
	}
	req.Header.Set("X-API-KEY", c.Token)
	resp, err := c.httpClient().Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// func (c *Client) error(StatusCode int, body io.Reader) error {
// 	var e APIError
// 	if err := json.NewDecoder(body).Decode(&e); err != nil {
// 		return &APIError{HTTPStatus: StatusCode}
// 	}
// 	e.HTTPStatus = StatusCode
// 	return &e
// }

func StructToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	if reflect.ValueOf(i).IsNil() {
		return
	}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		if !iVal.Field(i).IsZero() {
			values.Set(toSnakeCase(typ.Field(i).Name), fmt.Sprint(iVal.Field(i)))
		}
	}
	return
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
