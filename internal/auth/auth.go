package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPiKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")

	if val == "" {
		return "",errors.New("no authorisation info found")
	}

	vals := strings.Split(val," ")

	if len(vals) != 2{
		return "",errors.New("malformed auth header: "+ val)
	}

	if vals[0] != "ApiKey" {
		return "",errors.New("first part of auth header is malformed")
	}

	return vals[1],nil
}
