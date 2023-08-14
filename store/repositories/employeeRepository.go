package repositories

import (
	"EmployeeAPI/entity"
	"EmployeeAPI/store"
)

type EmployeeRepository struct {
	store *store.Store
}

func (repository *EmployeeRepository) Create(employee *entity.Employee) (*entity.Employee, error) {
	if err := repository.store.Db.QueryRow(
		"insert into employees (name, email) values ($1, $2) returning id",
		employee.Email, employee.Email).Scan(&employee.Id); err != nil {
		return nil, err
	}
	return employee, nil
}

func (repository *EmployeeRepository) FindByEmail(email string) (*entity.Employee, error) {
	result := entity.Employee{}
	if err := repository.store.Db.QueryRow(
		"select * from employees where email = $1",
		email).Scan(&result.Id, &result.Name, &result.Email); err != nil {
		return nil, err
	}
	return &result, nil
}
