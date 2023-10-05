package igboapi

import "fmt"

type APIError struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

// Error implements error.
func (err *APIError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("request failed with status code %d", err.HTTPStatus)
	}
	return fmt.Sprintf("StatusCode: %d, Code: %d, Message: %s", err.HTTPStatus, err.Code, err.Message)
}

var _ error = (*APIError)(nil)
// 2df2f3fa-5013-4602-93f7-02cab7be020c