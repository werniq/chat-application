package handlers

import (
	"bufio"
	"concurrency-chat/Logger"
	"concurrency-chat/models"
	"os"
	"strings"
)

func ReceiveMessage(m models.Message) bool {
	fi, err := os.Open("C:\\Users\\Oleksandr Matviienko\\projects\\golang\\concurrency-chat\\app\\handlers\\profanity-list.txt")
	if err != nil {
		Logger.ErrorLogger().Printf("error opening profanity-list.txt file: %v", err)
		return false
	}

	scanner := bufio.NewScanner(fi)
	var profanityList []string

	for scanner.Scan() {
		profanityList = append(profanityList, scanner.Text())
	}

	for i := 0; i < len(profanityList)-1; i++ {
		if strings.Contains(m.Content, profanityList[i]) {
			//var response struct {
			//	Error   bool   `json:"error"`
			//	Message string `json:"message"`
			//}
			//response.Error = true
			//response.Message = "Failed to send. Reveal profanity in content."
			//err := helpers.WriteJson(w, r, response, http.StatusBadRequest)
			if err != nil {
				Logger.ErrorLogger().Printf("error writing response: %v", err)
				return false
			}
		}
	}

	InsertIntoDatabase(m)
	return true
}
