package sib_db_hook

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)
import _ "github.com/lib/pq"

func TestGetNewArrivals(t *testing.T) {
	cnf, err := ReadConf("sample_cnf.yaml", "test")
	if err != nil {
		t.Fatalf("failure reading cnf file %v", err)
	}
	db, err := CreateConnection(cnf)
	if err != nil {
		t.Fatalf("failure creating connection %v", err)
	}
	defer db.Close()

	arrivalRequest := ArrivalQueryRequest{
		Date:      "2020-08-01",
		DateQuery: CreatedAt,
		Cursor:    0,
		Limit:     100,
	}
	arrivals, err := db.GetArrivals(arrivalRequest)
	if err != nil {
		t.Fatalf("failure executing query: %v", err)
	}

	if len(arrivals.Arrivals) == 0 {
		t.Error("want: non empty array, got: empty array")
	}

}

func TestGetArrivals_Updated(t *testing.T) {
	cnf, err := ReadConf("sample_cnf.yaml", "test")
	if err != nil {
		t.Fatalf("failure reading cnf file %v", err)
	}
	db, err := CreateConnection(cnf)
	if err != nil {
		t.Fatalf("failure creating connection %v", err)
	}
	defer db.Close()

	arrivalRequest := ArrivalQueryRequest{
		Date:      "2020-08-01",
		DateQuery: UpdateAt,
		Cursor:    0,
		Limit:     100,
	}
	arrivals, err := db.GetArrivals(arrivalRequest)
	if err != nil {
		t.Fatalf("failure executing query: %v", err)
	}

	if len(arrivals.Arrivals) == 0 {
		t.Error("want: non empty array, got: empty array")
	}

	if (arrivals.NextOffset) != 0 {
		t.Error("want: 0 offset, got non-0")
	}

}

type Location struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Index(vs []Location, t Location) int {
	for i, v := range vs {
		if v.Name == t.Name {
			return i
		}
	}
	return -1
}

func findDuplicates(l []Location) []Location {
	var encountered []Location

	for _, v := range l {
		if Index(encountered, v) < 0 {
			encountered = append(encountered, v)
		}
	}
	return encountered
}

func TestLocations(t *testing.T) {
	dat, _ := ioutil.ReadFile("locations.json")

	var locations []Location
	json.NewDecoder(bytes.NewReader(dat)).Decode(&locations)

	ds := findDuplicates(locations)
	t.Log(ds)

}
