package repository

import (
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (r *Repository) GetTasks() ([]model.Task, error) {
	tasks := make([]model.Task, 0, 10)
	res, err := r.db.Query("SELECT date, title, comment, repeat FROM scheduler ORDER BY date LIMIT 10")
	if err != nil {
		return tasks, fmt.Errorf("Repository.GetTasks select error: %w", err)
	}
	for res.Next() {
		var t model.Task
		err := res.Scan(&t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return tasks, fmt.Errorf("Repository.GetTasks scan error: %w", err)
		}
		if t.Title != "" {
			tasks = append(tasks, t)
		}
	}
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("Repository.GetTasks scan error: %w", err)
	}
	return tasks, nil
}
