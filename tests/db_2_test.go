package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type Task struct {
	ID      int64  `db:"id"`
	Date    string `db:"date"`
	Title   string `db:"title"`
	Comment string `db:"comment"`
	Repeat  string `db:"repeat"`
}

const (
	dbAdapterNamePg = "postgres"
)

func count(db *sqlx.DB) (int, error) {
	var count int
	return count, db.Get(&count, `SELECT count(id) FROM scheduler`)
}

func openDB(t *testing.T) *sqlx.DB {
	// dbfile := DBFile
	// envFile := os.Getenv("TODO_DBFILE")
	// if len(envFile) > 0 {
	// 	dbfile = envFile
	// }

	//db, err := sqlx.Connect("sqlite3", dbfile)

	appConfig, err := config.LoadAppConfig()
	if err != nil {
		panic(fmt.Sprintf("load config error: %v", err))
	}

	dataSourcePosgres := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbAdapterNamePg,
		appConfig.UserNamePG,
		appConfig.PasswordPG,
		appConfig.HostPG,
		appConfig.PortPG,
		appConfig.DbName)
	db, err := sqlx.Connect("postgres", dataSourcePosgres)
	assert.NoError(t, err)
	return db
}

func TestDB(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	before, err := count(db)
	assert.NoError(t, err)

	today := time.Now().Format(`20060102`)

	sqlStatement := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	res, err := db.Prepare(sqlStatement)
	// if err != nil {
	// 	log.Printf("Repository.AddTask insert error: %w\n", err)
	// }

	var id int
	err = res.QueryRow(today, "Todo", "Комментарий", "").Scan(&id)
	// if err != nil {
	// 	log.Printf("Repository.AddTask insert error: %w\n", err)
	// }

	// res, err := db.Exec(`INSERT INTO scheduler (date, title, comment, repeat)
	// VALUES (?, 'Todo', 'Комментарий', '')`, today)
	// assert.NoError(t, err)

	// id, err := res.LastInsertId()

	var task Task
	err = db.Get(&task, `SELECT * FROM scheduler WHERE id=$1`, id)
	//err = db.Get(&task, `SELECT * FROM scheduler WHERE id=?`, id)
	assert.NoError(t, err)
	assert.Equal(t, id, int(task.ID))
	assert.Equal(t, `Todo`, task.Title)
	assert.Equal(t, `Комментарий`, task.Comment)

	_, err = db.Exec(`DELETE FROM scheduler WHERE id = $1`, id)
	//_, err = db.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	assert.NoError(t, err)

	after, err := count(db)
	assert.NoError(t, err)

	assert.Equal(t, before, after)
}
