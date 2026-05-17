package task_test

import (
    "testing"
    "github.com/21v1u5/task-cli/internal/storage"
    "github.com/21v1u5/task-cli/internal/task"
)

func TestCreateTask(t *testing.T) {
    store := storage.NewMemoryStore()

    tsk := &task.Task{Title: "Estudar Go"}
    err := store.Create(tsk)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if tsk.ID == 0 {
        t.Error("expected ID to be set after Create")
    }
}

func TestCompleteTask(t *testing.T) {
    store := storage.NewMemoryStore()
    tsk := &task.Task{Title: "Tarefa"}
    store.Create(tsk)

    err := store.Complete(tsk.ID)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    found, _ := store.GetByID(tsk.ID)
    if found.Status != task.StatusDone {
        t.Errorf("expected status done, got %s", found.Status)
    }
}

func TestCompleteNotFound(t *testing.T) {
    store := storage.NewMemoryStore()
    err := store.Complete(999)
    if err == nil {
        t.Error("expected error for non-existent task")
    }
}

func TestDeleteTask(t *testing.T) {
    store := storage.NewMemoryStore()
    tsk := &task.Task{Title: "Deletar"}
    store.Create(tsk)

    store.Delete(tsk.ID)
    _, err := store.GetByID(tsk.ID)
    if err == nil {
        t.Error("expected error after deletion")
    }
}