package handlers

import (
	"concurrency-chat/Driver"
	"concurrency-chat/Logger"
	"concurrency-chat/models"
	"encoding/json"
	"net/http"
)

func InsertIntoDatabase(w http.ResponseWriter, r *http.Request) {
	db, err := Driver.OpenDB()
	if err != nil {
		Logger.ErrorLogger().Printf("error opening database connection: %v\n", err)
		return
	}
	mes := &models.Message{}
	err = json.NewDecoder(r.Body).Decode(&mes)

	err = db.InsertMessage(mes)
	if err != nil {
		Logger.ErrorLogger().Printf("error inserting message into database: %v\n", err)
		return
	}
	// c := make(chan *models.Message)

}
