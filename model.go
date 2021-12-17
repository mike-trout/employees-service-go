// model.go

package main

import (
	"errors"
	"strconv"

	"github.com/SoftwareAG/adabas-go-api/adabas"
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
	{
		PersonnelID: 11100102,
		FirstName:   "EDGAR",
		MiddleName:  "PETER",
		LastName:    "SCHINDLER",
	},
}

func getEmployees(start uint64, number uint64) ([]Employee, error) {
	// Create new connection handler to database
	connection, err := adabas.NewConnection("acj;target=1(adatcp://localhost:60001)")
	if err != nil {
		return nil, err
	}
	defer connection.Close()
	connection.Open()

	// To work on file 11 create corresponding read request
	request, err := connection.CreateFileReadRequest(11)
	if err != nil {
		return nil, err
	}
	// Define the result records content
	request.QueryFields("AA,AB")
	request.Limit = number

	// Read in the database using search query
	result, err := request.ReadLogicalWith("AA=[11100102:60021000]")
	if err != nil {
		return nil, err
	}

	// Add result to employees slice
	var employees []Employee
	var employee Employee
	for _, value := range result.Values {
		var aa, ac, ad, ae string
		value.Scan(&aa, &ac, &ad, &ae)
		p, err := strconv.Atoi(aa)
		if err != nil {
			return nil, err
		}
		employee.PersonnelID = p
		employee.FirstName = ac
		employee.MiddleName = ae
		employee.LastName = ad
		employees = append(employees, employee)
	}

	return employees, nil
}

func getEmployee(id int) (Employee, error) {
	for _, employee := range Employees {
		if employee.PersonnelID == id {
			return employee, nil
		}
	}

	return Employee{}, errors.New("Employee not found")
}
