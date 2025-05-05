package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	// title, content := getNoteData()
	// todoText := getTodoData()

	// userNote, err := note.New(title, content)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// todo, err := todo.New(todoText)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// todo.Display()
	// userNote.Display()
	// err = saveData(todo)

	// if err != nil {
	// 	fmt.Println("Saving the todo failed")
	// 	return
	// }

	// err = saveData(userNote)

	// if err != nil {
	// 	fmt.Println("Saving the note failed")
	// 	return
	// }

	// fmt.Println("Saving the note succeded")

	// outputData(todo)
	// outputData(userNote)
	printSomething(1)
	printSomething(1.5)

}

func outputData(data interface {
	Display()
	saver
}) {
	data.Display()
	saveData(data)
}

func printSomething(value any) {

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
	}

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)
	}
}

func saveData(entity saver) error {
	return entity.Save()
}

func getTodoData() string {
	text := getUserInput("Todo text:")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
