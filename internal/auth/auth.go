package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an API Key from the header of an HTTP request
//Example :
//Authorization: ApiKey  {insert API Key here}
func GetAPIKey(header http.Header) (string , error) {
	val:= header.Get("Authorization")

	if val == "" {
		return "", errors.New("no API Key provided")
	}
	vals:=strings.Split(val, " ")

	if len(vals) != 2  {
		return "", errors.New("malformed authorization header")
		
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authorization header")
	}

	return vals[1], nil
}