package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/celso-alexandre/learning-go-01/14-notes-and-todo/note"
	"github.com/celso-alexandre/learning-go-01/14-notes-and-todo/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Print()
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Print()
// 	Save() error
// }

func main() {
	title, content := getNotesData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(newNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoTitle := getUserInput("Enter todo title: ")
	newTodo, err := todo.New(todoTitle)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(newTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInput(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(input, "\n")
	text = strings.TrimSuffix(text, "\r") // for Windows
	return text
}

func getNotesData() (string, string) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")
	return title, content
}

func saveToFile(s saver) error {
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

func outputData(s outputtable) error {
	s.Print()
	return s.Save()
}
