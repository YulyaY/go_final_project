package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

const (
	limit               = 20
	FormatDateForSearch = "02.01.2006"
)

func (r *Repository) GetTasks(search string) ([]model.Task, error) {
	tasks := make([]model.Task, 0, 10)
	var res *sql.Rows
	var err error
	if search != "" {
		searchDate, errDateParse := time.Parse(FormatDateForSearch, search)
		var date string
		if errDateParse != nil {
			search = fmt.Sprintf("%%%s%%", search)
			res, err = r.db.Query("SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit",
				sql.Named("search", search),
				sql.Named("limit", limit))
		} else {
			date = domain.Format(searchDate)
			res, err = r.db.Query("SELECT * FROM scheduler WHERE date = :date ORDER BY date LIMIT :limit",
				sql.Named("date", date),
				sql.Named("limit", limit))
		}
	} else {
		res, err = r.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit))
	}
	if err != nil {
		return tasks, fmt.Errorf("Repository.GetTasks select error: %w", err)
	}
	for res.Next() {
		var t model.Task
		err := res.Scan(&t.Id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
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
