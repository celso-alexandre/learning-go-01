package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/celso-alexandre/learning-go-01/14-notes-and-todo/note"
	"github.com/celso-alexandre/learning-go-01/14-notes-and-todo/todo"
)

func main() {
	title, content := getNotesData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	newNote.Print()
	err = newNote.Save()
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
	newTodo.Print()
	err = newTodo.Save()
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
	// var input string
	// fmt.Scanln(&input)
	// return input
}

func getNotesData() (string, string) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")
	return title, content
}
