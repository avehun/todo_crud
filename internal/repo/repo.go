package repo

import (
	"context"

	"github.com/avehun/todo_crud/internal/model"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) *Repo {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return &Repo{db: db}
}
func (repo *Repo) GetTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	rows, _ := repo.db.Query(context.Background(), "select * from tasks")

	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Created_at,
			&task.Updated_at)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *Repo) AddTask(task model.Task) error {
	_, err := repo.db.Exec(context.Background(), "INSERT INTO tasks(title, description, status) VALUES($1,$2,$3)", task.Title, task.Description, task.Status)
	return err
}
func (repo *Repo) UpdateTask(task model.Task) error {
	_, err := repo.db.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, task.Updated_at, task.Id)
	return err
}
func (repo *Repo) DeleteTask(id int) error {
	_, err := repo.db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	return err
}
