package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID          int       `json:"id"`
	MessageID   string    `json:"message_id"`
	Username    string    `json:"username"`
	AuthorId    string    `json:"author_id"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"name"`
	Email        string    `json:"email"`
	Password     []byte    `json:"password"`
	MessagesSent int       `json:"messages_sent"`
	CreatedTime  time.Time `json:"created_time"`
}

type DatabaseModel struct {
	DB *sql.DB
}

func (db *DatabaseModel) InsertMessage(m Message) error {
	stmt := `
		INSERT INTO 
		    messages
		    (message_id, content, created_time, author_username, author_id) 
		VALUES
		    ($1, $2, $3, $4)
		    			`

	_, err := db.DB.Exec(stmt, m.ID, m.Content, m.CreatedTime, m.Username, m.AuthorId)
	if err != nil {
		return err
	}
	return nil
}
