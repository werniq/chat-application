package handlers

import (
	"bytes"
	"concurrency-chat/Logger"
	"concurrency-chat/models"
	"encoding/json"
	"net/http"
)

func SendMessage(w http.ResponseWriter, r *http.Request) *models.Message {
	mes := &models.Message{}
	err := json.NewDecoder(r.Body).Decode(&mes)
	if err != nil {
		Logger.ErrorLogger().Printf("error decoding request body: %v", err)
		return nil
	}

	out, err := json.Marshal(mes)
	if err != nil {
		Logger.ErrorLogger().Printf("error marshalling data: %v\n", err)
		return nil
	}

	req, err := http.NewRequest("POST", "http://", bytes.NewBuffer(out))
	if err != nil {
		Logger.ErrorLogger().Printf("error creating request: %v\n", err)
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return mes
}
