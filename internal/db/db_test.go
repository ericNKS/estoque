package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load("../../.env")
	_, err := Connection()
	if err != nil {
		t.Fatal(err)
	}

}
