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
	title     string
	content   string
	createdAt time.Time
}

type publicNote struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content \n\n %v\n\n", note.title, note.content)
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(publicNote{Title: note.title, Content: note.content, CreatedAt: note.createdAt})

	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(json), 0644)
}
