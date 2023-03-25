package helpers

import (
	"concurrency-chat/Logger"
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, r *http.Request, data interface{}, status int) error {
	out, err := json.Marshal(data)
	if err != nil {
		Logger.ErrorLogger().Printf("error marshalling data: %v", err)
		return err
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	return nil
}
