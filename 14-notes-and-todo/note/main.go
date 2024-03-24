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
	filename  string    // not exported
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}
	filename := strings.ReplaceAll(title, " ", "-")
	filename = strings.ToLower(filename)
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		filename:  filename,
	}, nil
}

func (n *Note) Print() {
	fmt.Println()
	fmt.Printf("Your note titled %v has been created at %v.\nContent: %v\n", n.Title, n.CreatedAt, n.Content)
}

func (n Note) ToJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Note) Save() error {
	content, err := n.ToJSON()
	if err != nil {
		return err
	}
	return os.WriteFile(n.filename+".json", []byte(content), 0644)
}
