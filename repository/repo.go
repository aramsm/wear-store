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

func logError(e error) {
	fmt.Println("Error: ", e)
	os.Exit(1)
}
