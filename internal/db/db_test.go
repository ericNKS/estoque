package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	db, err := Connection()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
}
