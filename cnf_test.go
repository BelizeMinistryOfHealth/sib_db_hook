package sib_db_hook

import "testing"

func TestReadConf_SuccessfullyParsesConfig(t *testing.T) {
	cnf, err := ReadConf("sample_cnf.yaml", "test")
	if err != nil {
		t.Fatalf("failed to read cnf file: %v", err)
	}

	if cnf.DbDatabase == "" {
		t.Fatalf("want: %s, got: %s", "database", cnf.DbDatabase)
	}
}
