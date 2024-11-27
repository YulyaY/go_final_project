package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteTask(id int) error {
	sqlResult, err := r.db.Exec("DELETE FROM scheduler WHERE id = :id",
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("Repository.DeleteTask delete error: %w", err)
	}

	affectedCnt, err := sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("Repository.DeleteTask delete error: %w", err)
	}
	if affectedCnt == 0 {
		return errRecordDoesNotExists
	}
	fmt.Println(affectedCnt)
	return nil
}
