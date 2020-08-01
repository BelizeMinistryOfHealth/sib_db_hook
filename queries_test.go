package sib_db_hook

import (
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

	arrivals, err := GetArrivals("2020-08-01", CreatedAt, db)
	if err != nil {
		t.Fatalf("failure executing query: %v", err)
	}

	if len(arrivals) == 0 {
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

	arrivals, err := GetArrivals("2020-08-01", UpdateAt, db)
	if err != nil {
		t.Fatalf("failure executing query: %v", err)
	}

	if len(arrivals) == 0 {
		t.Error("want: non empty array, got: empty array")
	}

}
