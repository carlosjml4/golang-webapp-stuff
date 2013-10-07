package vulcan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (r *HttpError) Error() string {
	return fmt.Sprintf(
		"HttpError(code=%d, %s, %s)", r.StatusCode, r.Status, r.Body)
}

func NewHttpError(statusCode int) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Body:       []byte(http.StatusText(statusCode))}
}

func TooManyRequestsError(retrySeconds int) *HttpError {

	encodedError, err := json.Marshal(map[string]interface{}{
		"error":         "Too Many Requests",
		"retry-seconds": retrySeconds,
	})

	if err != nil {
		// something terrible just happened
		// if json encoder fails, I don't know what to do :-/
		panic(err)
	}

	return &HttpError{
		StatusCode: 429, //RFC 6585
		Status:     "Too Many Requests",
		Body:       encodedError}
}

// We somehow failed to authenticate the request
type AuthError string

func (f AuthError) Error() string {
	return string(f)
}
