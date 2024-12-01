package usernote

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type NoteType struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note NoteType) DisplayTitle() {
	fmt.Println(note.Title)
}

func (note NoteType) DisplayContent() {
	fmt.Println(note.Content)
}

func (note NoteType) Display() {
	fmt.Printf("Your note title %v has the following content: \n %v \n", note.Title, note.Content)
}

func (note NoteType) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	data, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func New(title, content string) (NoteType, error) {
	if title == "" || content == "" {
		return NoteType{}, errors.New("field should not be empty")
	}
	return NoteType{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
