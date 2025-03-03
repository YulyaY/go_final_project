package repository

import (
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (r *Repository) AddTask(t model.Task) (int, error) {
	sqlStatement := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	res, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return 0, fmt.Errorf("Repository.AddTask insert error: %w", err)
	}

	var id int
	err = res.QueryRow(t.Date, t.Title, t.Comment, t.Repeat).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Repository.AddTask insert error: %w", err)
	}

	// res, err := r.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
	// 	sql.Named("date", t.Date),
	// 	sql.Named("title", t.Title),
	// 	sql.Named("comment", t.Comment),
	// 	sql.Named("repeat", t.Repeat))

	// if err != nil {
	// 	return 0, fmt.Errorf("Repository.AddTask insert error: %w", err)
	// }
	// idLast, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("Repository.AddTask insert error: %w", err)
	// }
	return id, nil
}
