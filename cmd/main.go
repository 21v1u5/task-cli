package main

import (
	"fmt"
	"os"
	"github.com/21v1u5/task-cli/internal/cli"
	"github.com/21v1u5/task-cli/internal/storage"
)

func main() {
	store, err := storage.NewSQLiteStore("tasks.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening database: %v", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}

	handler := cli.NewHandler(store)

	switch os.Args[1] {
	case "add":
		handler.Add(os.Args[2:])
	case "list":
		handler.List(os.Args[2:])
	case "done":
		handler.Complete(os.Args[2:])
	case "delete":
		handler.Delete(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s", os.Args[1])
		cli.PrintUsage()
		os.Exit(1)
	}
}