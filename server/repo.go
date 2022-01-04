package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndParseData() []Store {
	bs, readErr := ioutil.ReadFile("acme-stores.json")

	if readErr != nil {
		fmt.Println("Error: ", readErr)
		os.Exit(1)
	}

	var stores []Store
	jsonErr := json.Unmarshal(bs, &stores)

	if jsonErr != nil {
		fmt.Println("Error: ", jsonErr)
		os.Exit(1)
	}

	return stores
}
