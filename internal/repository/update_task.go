package repository

import (
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

func (r *Repository) UpdateTask(t model.Task) error {
	sqlResult, err := r.db.Exec("UPDATE scheduler SET date = $1, title = $2, comment = $3, repeat = $4  WHERE id = $5",
		t.Date, t.Title, t.Comment, t.Repeat, t.Id)

	// sqlResult, err := r.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat  WHERE id = :id",
	// 	sql.Named("date", t.Date),
	// 	sql.Named("title", t.Title),
	// 	sql.Named("comment", t.Comment),
	// 	sql.Named("repeat", t.Repeat),
	// 	sql.Named("id", t.Id))
	if err != nil {
		return fmt.Errorf("Repository.UpdateTask update error: %w", err)
	}

	affectedCnt, err := sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("Repository.UpdateTask update error: %w", err)
	}
	if affectedCnt == 0 {
		return service.ErrRecordDoesNotExists
	}
	return nil
}
