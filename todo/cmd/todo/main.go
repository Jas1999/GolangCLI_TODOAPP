package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/Jas1999/GolangCLI_TODOAPP/tree/main"
)

const (
	todoFile = "todos.json"
)

func main() {
	fmt.Println("hello TODO cli")

	add := flag.Bool("add", false, "Add a new todo")
	list := flag.Bool("list", false, "list tasks")
	complete := flag.Int("complete", 0, "mark task as complete")
	delete := flag.Int("delete, ", 0, "delete task")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:

		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid CMD")
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r) // this pipe case : echo "test" | ./todo -add
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", errors.New("empty is not allowed")
	}
	return scanner.Text(), nil
}
