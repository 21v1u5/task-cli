package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"github.com/21v1u5/task-cli/internal/storage"
	"github.com/21v1u5/task-cli/internal/task"
)

type Handler struct{ store storage.Store }

func NewHandler(store storage.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Add(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	desc := fs.String("desc", "", "task description")
	fs.Parse(args)

	if fs.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "usage: task-cli add: %s",fs)
	}
	return nil
}