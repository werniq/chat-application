package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID         int       `json:"id"`
	MessageID  int       `json:"messageID"`
	Content    string    `json:"content"`
	AuthorID   int       `json:"author_id"`
	Author     *User     `json:"author"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
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

func (db *DatabaseModel) InsertMessage(m *Message) error {
	stmt := `
		INSERT INTO 
		    messages
		    (message_id, content, created_time, modified, author_id) 
		VALUES
		    ($1, $2, $3, $4, $5)
		    `

	_, err := db.DB.Query(stmt, m.MessageID, m.Content, m.CreatedAt, m.ModifiedAt, m.AuthorID)
	if err != nil {
		return err
	}
	return nil
}
