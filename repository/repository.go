package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndParseData() []Store {
	bs, readErr := ioutil.ReadFile("/home/aram.menocci/workspace/wear-store/acme-stores.json")

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

func GetStoresByBrand(data []Store, brand string) []Store {
	var result []Store

	for i := range data {
		if data[i].BrandLabel == brand {
			result = append(result, data[i])
		}
	}

	return result
}

func GetStore(data []Store, id string) Store {
	var store Store

	for i := range data {
		if data[i].ID == id {
			store = data[i]

			return store
		}
	}

	return store
}

func GetEmployees(data []Store) []Employee {
	var employees = make([]Employee, 0)

	for i := range data {
		employees = append(employees, data[i].Employees...)
	}

	return employees
}

func GetEmployee(data []Store, id string) Employee {
	var employee Employee

	for s := range data {
		for e := range data[s].Employees {
			if data[s].Employees[e].ID == id {
				employee = data[s].Employees[e]

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
