package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content are required")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display() {
	fmt.Printf("\nYour note titled '%v' has the following content:\n%v\n\n", note.Title, note.Content)
}

func (note *Note) Save() error {
	fileName := strings.ToLower(note.Title) + ".json"
	fileName = strings.ReplaceAll(fileName, " ", "_")

	json, err := json.Marshal(&note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
