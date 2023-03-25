package helpers

import (
	"concurrency-chat/Logger"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ReadJson reads data from request, and parses it into data
func ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// i want to handle maximum 1 megabyte of data
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		Logger.ErrorLogger().Printf("error decoding request body: %v", err)
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain only one json value")
	}
	return nil
}
