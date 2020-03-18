// model.go

package main

import (
	"errors"
)

// Employee - struct modelling an employee record
type Employee struct {
	PersonnelID int    `json:"personnelId"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastname"`
}

// Employees - global var containing all employees
var Employees = []Employee{
	Employee{
		PersonnelID: 11100102,
		FirstName:   "EDGAR",
		MiddleName:  "PETER",
		LastName:    "SCHINDLER",
	},
}

func getEmployees() ([]Employee, error) {
	return Employees, nil
}

func getEmployee(id int) (Employee, error) {
	for _, employee := range Employees {
		if employee.PersonnelID == id {
			return employee, nil
		}
	}

	return Employee{}, errors.New("Employee not found")
}
