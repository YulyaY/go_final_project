package repository

import (
	"database/sql"
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (r *Repository) GetTask(id int) (model.Task, error) {
	var t model.Task
	res := r.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))
	err := res.Scan(&t.Id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return t, fmt.Errorf("Repository.GetTask select error: %w", err)
	}
	return t, nil
}
