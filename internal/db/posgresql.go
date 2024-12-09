package db

import (
	"database/sql"
	"fmt"

	"github.com/YulyaY/go_final_project.git/internal/config"
	_ "github.com/lib/pq"
)

const (
	dbAdapterNamePg = "postgres"
)

func NewPosgres(appCfg config.AppConfig) (*sql.DB, error) {
	// _, err := os.Stat(dbName)
	// if err != nil {
	// 	return nil, err
	// }

	dataSourcePosgres := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbAdapterNamePg,
		appCfg.UserNamePG,
		appCfg.PasswordPG,
		appCfg.HostPG,
		appCfg.PortPG,
		appCfg.DbName)

	db, err := sql.Open(dbAdapterNamePg, dataSourcePosgres)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Posgresql db is open")

	return db, nil
}

func CreateDbPostgres(appCfg config.AppConfig) (*sql.DB, error) {
	dataSourcePosgres := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbAdapterNamePg,
		appCfg.UserNamePG,
		appCfg.PasswordPG,
		appCfg.HostPG,
		appCfg.PortPG,
		appCfg.DbName)

	db, err := sql.Open(dbAdapterNamePg, dataSourcePosgres)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Posgresql db is open")

	return db, nil
}
