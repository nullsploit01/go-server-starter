package response

import (
	"encoding/json"
	"net/http"

	m "github.com/nullsploit01/server/models"
)

func JSON[T any](w http.ResponseWriter, status int, data T) error {
	response := m.ResponseBody[T]{
		Error: false,
		Data:  data,
	}

	return JSONWithHeaders(w, status, response, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}
