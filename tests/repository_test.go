package tests

import (
	"testing"

	"github.com/aramsm/wear-store/repository"
)

func TestGetStore(t *testing.T) {
	id := "76"
	data := repository.ReadAndParseData()
	result := repository.GetStore(data, id)

	if result.ID != id {
		t.Errorf("Expected to found store with id '%v', found '%v' instead", id, result.ID)
	}
}

func TestGetStoresByBrand(t *testing.T) {
	brand := "acme-inc"
	data := repository.ReadAndParseData()
	results := repository.GetStoresByBrand(data, brand)

	if len(results) == 0 {
		t.Errorf("No store found with brand '%v'", brand)
	}

	for i := range results {
		if results[i].BrandLabel != brand {
			t.Errorf("Expected to all stores be from brand '%v'", brand)
		}
	}
}

func TestGetEmployee(t *testing.T) {
	id := "1"
	data := repository.ReadAndParseData()
	result := repository.GetEmployee(data, id)

	if result.ID != id {
		t.Errorf("Expected to found employee with id '%v', found '%v' instead", id, result.ID)
	}
}
