import storage

import (
	"database/sql"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/21v1u5/task-cli/internal/task"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStorage(path string) (*SQLiteStore, error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	store := &SQLiteStore{db: db}
	return store, store.migrate()
}

func (s *SQLiteStore) migrate() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			title       TEXT NOT NULL,
			description TEXT,
			status      TEXT NOT NULL DEFAULT 'pending',
			created_at  DATETIME NOT NULL,
			updated_at  DATETIME NOT NULL
		)
	`)
	return err
}

func (s *SQLiteStore) Create(t *task.Task) error {
	now := time.Now()
	result, err := s.db.Exec(
		`INSERT INTO tasks (title, description, status, created_at, update_at)
		VALUES (?, ?, ?, ?, ?)`,
		t.Title, t.Description, task.StatusPending, now, now,
	)
	if err != nil {
		return err
	}
	t.ID, _ = result.LastInsertId()
	t.CreatedAt = now
	t.UpdatedAt = now
	return nil
}

func (s *SQLiteStore) List() ([]task.Task, error){
	rows, err := s.db.Query(
		`SELECT id, title, description, status, created_at, updated_at
		FROM tasks ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		err := rows.Scan(
			&t.ID, &t.Title, &t.Description,
			&t.Status, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}