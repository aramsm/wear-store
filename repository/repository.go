package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndParseData() []Store {
	bs, readErr := ioutil.ReadFile("acme-stores.json")

	if readErr != nil {
		logError(readErr)
	}

	var stores []Store
	jsonErr := json.Unmarshal(bs, &stores)

	if jsonErr != nil {
		logError(jsonErr)
	}

	return stores
}

func GetStore(id string) Store {
	var store Store
	stores := ReadAndParseData()

	for _, s := range stores {
		if s.ID == id {
			store = s

			return store
		}
	}

	return store
}

func GetEmployees() []Employee {
	var employees = make([]Employee, 0)
	stores := ReadAndParseData()

	for _, s := range stores {
		employees = append(employees, s.Employees...)
	}

	return employees
}

func GetEmployee(id string) Employee {
	var employee Employee
	stores := ReadAndParseData()

	for _, s := range stores {
		for _, e := range s.Employees {
			if e.ID == id {
				employee = e

				return employee
			}
		}
	}

	return employee
}

func logError(e error) {
	fmt.Println("Error: ", e)
	os.Exit(1)
}
