package repository

import (
	"database/sql"
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (r *Repository) GetTasks(taskFilter model.GetTaskFilter, limit int) ([]model.Task, error) {
	tasks := make([]model.Task, 0, 10)
	var res *sql.Rows
	var err error
	if taskFilter.DateFilter != nil {
		res, err = r.db.Query("SELECT * FROM scheduler WHERE date = $1 ORDER BY date LIMIT $2",
			*taskFilter.DateFilter, limit)

		// res, err = r.db.Query("SELECT * FROM scheduler WHERE date = :date ORDER BY date LIMIT :limit",
		// 	sql.Named("date", *taskFilter.DateFilter),
		// 	sql.Named("limit", limit))
	} else if taskFilter.TitleFilter != nil {
		res, err = r.db.Query("SELECT * FROM scheduler WHERE title LIKE $1 OR comment LIKE $1 ORDER BY date LIMIT $2",
			*taskFilter.TitleFilter, limit)

		// res, err = r.db.Query("SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit",
		// 	sql.Named("search", *taskFilter.TitleFilter),
		// 	sql.Named("limit", limit))
	} else {
		res, err = r.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT $1", limit)

		//res, err = r.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit))
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
