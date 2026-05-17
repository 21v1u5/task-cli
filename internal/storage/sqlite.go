import storage

import (
	"database/sql"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/21v1u5/task-cli/internal/task"
)

type Store interface {
	Create(t *task.Task) error
	List() ([]task.Task, error)
	GetByID(id int64) (*task.Task, error)
	Complete(id int64) error
	Delete(id int64) error
}

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

func (s *SQLiteStore) GetByID(id int64) (*task.Task, err) {
	var t task.Task
	err := s.db.QueryRow(
		`SELECT id, title, description, status, created_at, updated_at
		FROM tasks Where id = ?`, id,
	).Scan(&t.ID, &t.Title, &t.Description,
		   &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %d not found", id)
	}
	return &t, err
}

func (s *SQLiteStore) Complete(id int64) error {
	result, err := s.db.Exec(
		`UPDATE tasks SET status = ?, updated_at = ?
		WHERE id = ? AND status = 'pending'`,
		task.StatusDone, time.Now(), id,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("task %d not found or already done", id)
	}
	return nil
}

func (s *SQLiteStore) Delete(id int64) error {
	result, err := s.db.Exec(
		"DELETE FROM tasks WHERE id = ?", id,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("task %d not found", id)
	}
	return nil
}