package store

import (
	"EmployeeAPI/store/repositories"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type Store struct {
	config             *Config
	Db                 *sql.DB
	employeeRepository *repositories.EmployeeRepository
	taskRepository     *repositories.TaskRepository
}

func New(config *Config) *Store {
	return &Store{config: config}
}

func (store *Store) Open() error {
	db, err := sql.Open(driverName, store.config.DataBaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func (store *Store) Closer() error {
	err := store.Db.Close()
	if err != nil {
		return err
	}
	return nil
}
