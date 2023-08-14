package repositories

import (
	"EmployeeAPI/entity"
	"EmployeeAPI/store"
)

type TaskRepository struct {
	store *store.Store
}

func (repository *TaskRepository) Create(task *entity.Task) (*entity.Task, error) {
	if err := repository.store.Db.QueryRow(
		"insert into task (name, description) values ($1, $2) returning id",
		task.Name, task.Description).Scan(&task.Id); err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *TaskRepository) FindByName(name string) (*entity.Task, error) {
	result := entity.Task{}
	if err := repository.store.Db.QueryRow(
		"select * from task where name = $1",
		name).Scan(&result.Id, &result.Name, &result.Description); err != nil {
		return nil, err
	}
	return &result, nil
}
