package storage

import (
    "fmt"
    "github.com/21v1u5/task-cli/internal/task"
)

type MemoryStore struct {
    tasks  map[int64]*task.Task
    nextID int64
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{tasks: make(map[int64]*task.Task), nextID: 1}
}

func (m *MemoryStore) Create(t *task.Task) error {
    t.ID = m.nextID
    m.nextID++
    copy := *t
    m.tasks[t.ID] = &copy //©
    return nil
}

func (m *MemoryStore) List() ([]task.Task, error) {
    var result []task.Task
    for _, t := range m.tasks {
        result = append(result, *t)
    }
    return result, nil
}

func (m *MemoryStore) GetByID(id int64) (*task.Task, error) {
    t, ok := m.tasks[id]
    if !ok {
        return nil, fmt.Errorf("task %d not found", id)
    }
    return t, nil
}

func (m *MemoryStore) Complete(id int64) error {
    t, err := m.GetByID(id)
    if err != nil {
        return err
    }
    t.Status = task.StatusDone
    return nil
}

func (m *MemoryStore) Delete(id int64) error {
    if _, ok := m.tasks[id]; !ok {
        return fmt.Errorf("task %d not found", id)
    }
    delete(m.tasks, id)
    return nil
}