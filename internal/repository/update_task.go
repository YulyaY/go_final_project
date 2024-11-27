package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

var errRecordDoesNotExists error = errors.New("record does not exists")

func (r *Repository) UpdateTask(t model.Task) error {
	sqlResult, err := r.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat  WHERE id = :id",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.Id))
	if err != nil {
		return fmt.Errorf("Repository.UpdateTask update error: %w", err)
	}

	affectedCnt, err := sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("Repository.UpdateTask update error: %w", err)
	}
	if affectedCnt == 0 {
		return errRecordDoesNotExists
	}
	return nil
}
