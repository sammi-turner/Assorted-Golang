package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	ID   uint `gorm:"primary_key"`
	Task string
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "todos.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
	displayMenu()
}

func clearConsole() {
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Flushes stdout.
func flushStdout() {
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()
}

// Prompts the user with a given string and returns their input.
func userInput(s string) string {
	fmt.Print(s)
	flushStdout()
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

func displayMenu() {
	for {
		readTodos()
		input := userInput("\n\nCreate (c), Delete (d) or Exit (e)? ")
		switch input {
		case "C", "c":
			createTodo()
		case "D", "d":
			deleteTodo()
		case "E", "e":
			fmt.Println()
			os.Exit(0)
		default:
			// Do nothing
		}
	}
}

func createTodo() {
	reader := bufio.NewReader(os.Stdin)
	var task string
	var id uint
	fmt.Print("Enter task ID: ")
	input, _ := reader.ReadString('\n')
	idUint64, err := strconv.ParseUint(strings.TrimSpace(input), 10, 32)
	if err != nil {
		displayMenu()
	}
	id = uint(idUint64)
	fmt.Print("Task name? ")
	task, _ = reader.ReadString('\n')
	task = strings.TrimSpace(task)

	todo := Todo{ID: id, Task: task}
	db.Save(&todo)
}

func readTodos() {
	clearConsole()
	fmt.Print("\nTUI TODOS\n")
	var todos []Todo
	db.Find(&todos)
	for _, todo := range todos {
		fmt.Printf("\n%d. %s", todo.ID, todo.Task)
	}
}

func deleteTodo() {
	reader := bufio.NewReader(os.Stdin)
	var id uint

	fmt.Print("ID to delete? ")
	input, _ := reader.ReadString('\n')
	idUint64, err := strconv.ParseUint(strings.TrimSpace(input), 10, 32)
	if err != nil {
		displayMenu()
	}
	id = uint(idUint64)

	todo := Todo{}
	db.First(&todo, id)
	db.Delete(&todo)
}
