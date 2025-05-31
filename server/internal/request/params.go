package request

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nullsploit01/server/internal/customError"
)

func GetURLParamInt(r *http.Request, key string) (int, error) {
	value, err := GetURLParamString(r, key)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func GetURLParamString(r *http.Request, key string) (string, error) {
	value := chi.URLParam(r, key)

	if value == "" {
		return "", fmt.Errorf("missing %s parameter", key)
	}

	return value, nil
}

func GetURLParamBool(r *http.Request, key string) (bool, error) {
	value, err := GetURLParamString(r, key)
	if err != nil {
		return false, err
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}

func GetURLQueryParamInt(r *http.Request, key string, required bool) (int, error) {
	value, err := GetURLQueryParamString(r, key, required)
	if err != nil {
		return 0, err
	}

	if value == "" {
		if required {
			return 0, customError.ValidationError{Field: key, Message: "missing query parameter"}
		}
		return 0, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func GetURLQueryParamString(r *http.Request, key string, required bool) (string, error) {
	value := r.URL.Query().Get(key)

	if value == "" && required {
		return "", fmt.Errorf("missing %s parameter", key)
	}

	return value, nil
}

func GetURLQueryParamBool(r *http.Request, key string, required bool) (bool, error) {
	value, err := GetURLQueryParamString(r, key, required)
	if err != nil {
		return false, err
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}
