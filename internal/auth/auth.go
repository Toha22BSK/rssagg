package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extacts an API key from
// the headers of an HTTP request.
// Example:
// Authorization: ApiKey {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Authorization header")
	}
	return vals[1], nil
}
