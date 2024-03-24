package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title    string `json:"title"`
	filename string
}

func New(title string) (*Note, error) {
	if title == "" {
		return nil, errors.New("title and content are required")
	}
	filename := strings.ReplaceAll(title, " ", "-")
	filename = strings.ToLower(filename)
	return &Note{
		Title:    title,
		filename: filename,
	}, nil
}

func (n *Note) Print() {
	fmt.Println()
	fmt.Printf("Your note titled %v has been created.", n.Title)
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
