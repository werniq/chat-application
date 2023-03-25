package handlers

import (
	"concurrency-chat/Driver"
	"concurrency-chat/Logger"
	"concurrency-chat/models"
)

func InsertIntoDatabase(m models.Message) {
	db, err := Driver.OpenDB()
	if err != nil {
		Logger.ErrorLogger().Printf("error opening database connection: %v\n", err)
		return
	}

	err = db.InsertMessage(m)
	if err != nil {
		Logger.ErrorLogger().Printf("error inserting message into database: %v\n", err)
		return
	}
	// c := make(chan *models.Message)

}
