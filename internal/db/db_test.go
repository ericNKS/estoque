package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := Connection()
	if err != nil {
		t.Fatal(err)
	}

}
