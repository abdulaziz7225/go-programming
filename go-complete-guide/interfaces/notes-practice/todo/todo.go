package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("invalid todo input")
	}
	return &Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}

func (note *Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(&note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
