package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jrgmonsalve/go-labs/udemy/6s_interfaces/note"
	"github.com/jrgmonsalve/go-labs/udemy/6s_interfaces/todo"
)

type Saver interface {
	Save() error
}

type outPutable interface {
	Display()
	Saver
}

func main() {
	printSomeThing("Hello world")
	printSomeThing(123)
	printSomeThing(123.123)

	result := sum(1, 2)
	fmt.Println("Sum:", result)

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outPutData(userTodo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	outPutData(userNote)

}

func outPutData(o outPutable) error {
	o.Display()
	err := saveData(o)
	if err != nil {
		return err
	}
	return nil
}

func saveData(s Saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}
	fmt.Println("Data saved successfully.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

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

func printSomeThing(data interface{}) {

	//option 1
	switch data.(type) {
	case string:
		fmt.Println("String:", data)
	case int:
		fmt.Println("Int:", data)
	case float64:
		fmt.Println("Float64:", data)
	}
	fmt.Println(data)

	//option 2
	typeValue, ok := data.(int)
	if ok {
		fmt.Println("Int:", typeValue)
	}

	//option 3
	typeValue3 := fmt.Sprintf("%v", data)
	fmt.Println("Type 3:", typeValue3)
}

func sum[T int | float64 | string](a, b T) T {
	return a + b
}
