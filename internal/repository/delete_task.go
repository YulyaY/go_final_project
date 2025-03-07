package repository

import (
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

func (r *Repository) DeleteTask(id int) error {

	sqlResult, err := r.db.Exec("DELETE FROM scheduler WHERE id = $1", id)
	// sqlResult, err := r.db.Exec("DELETE FROM scheduler WHERE id = :id",
	// 	sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Repository.DeleteTask delete error: %w", err)
	}

	affectedCnt, err := sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("Repository.DeleteTask delete error: %w", err)
	}
	if affectedCnt == 0 {
		return service.ErrRecordDoesNotExists
	}
	return nil
}
