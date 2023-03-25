package handlers

import (
	"bufio"
	"concurrency-chat/Logger"
	"concurrency-chat/app/handlers/helpers"
	"net/http"
	"os"
	"strings"
)

func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	var payload struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		AuthorID string `json:"author_id"`
		Author   string `json:"author"`
	}

	err := helpers.ReadJson(w, r, &payload)
	if err != nil {
		Logger.ErrorLogger().Printf("error reading json: %v", err)
	}
	fi, err := os.Open("C:\\Users\\Oleksandr Matviienko\\projects\\golang\\concurrency-chat\\app\\handlers\\profanity-list.txt")
	if err != nil {
		Logger.ErrorLogger().Printf("error opening profanity-list.txt file: %v", err)
		return
	}

	scanner := bufio.NewScanner(fi)
	var profanityList []string

	for scanner.Scan() {
		profanityList = append(profanityList, scanner.Text())
	}

	for i := 0; i < len(profanityList)-1; i++ {
		if strings.Contains(payload.Content, profanityList[i]) {
			var response struct {
				Error   bool   `json:"error"`
				Message string `json:"message"`
			}
			response.Error = true
			response.Message = "Failed to send. Reveal profanity in contenta."
			err := helpers.WriteJson(w, r, response, http.StatusBadRequest)
			if err != nil {
				Logger.ErrorLogger().Printf("error writing response: %v", err)
				return
			}
		}
	}

	InsertIntoDatabase(w, r)
}
