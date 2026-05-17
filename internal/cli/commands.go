package cli

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "github.com/seu-usuario/task-cli/internal/storage"
    "github.com/seu-usuario/task-cli/internal/task"
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
        fmt.Fprintln(os.Stderr, "usage: task-cli add <title> [-desc <description>]")
        os.Exit(1)
    }

    t := &task.Task{
        Title:       fs.Arg(0),
        Description: *desc,
    }
    if err := h.store.Create(t); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Task #%d created: %s\n", t.ID, t.Title)
}

func (h *Handler) List(args []string) {
    fs := flag.NewFlagSet("list", flag.ExitOnError)
    all := fs.Bool("all", false, "show completed tasks too")
    fs.Parse(args)

    tasks, err := h.store.List()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("%-4s %-6s %s\n", "ID", "STATUS", "TITLE")
    fmt.Println("─────────────────────────────────")
    for _, t := range tasks {
        if !*all && t.Status == "done" {
            continue
        }
        status := "[ ]"
        if t.Status == "done" {
            status = "[x]"
        }
        fmt.Printf("%-4d %-6s %s\n", t.ID, status, t.Title)
    }
}

func (h *Handler) Complete(args []string) {
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr, "usage: task-cli done <id>")
        os.Exit(1)
    }
    id, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
        fmt.Fprintln(os.Stderr, "error: id must be a number")
        os.Exit(1)
    }
    if err := h.store.Complete(id); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Task #%d marked as done\n", id)
}

func (h *Handler) Delete(args []string) {
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr, "usage: task-cli delete <id>")
        os.Exit(1)
    }
    id, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
        fmt.Fprintln(os.Stderr, "error: id must be a number")
        os.Exit(1)
    }
    if err := h.store.Delete(id); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Task #%d deleted\n", id)
}

func PrintUsage() {
    fmt.Println("usage: task-cli <command> [options]")
    fmt.Println("commands: add, list, done, delete")
}