package db

import "testing"

func TestInitDB(t *testing.T) {
	db := InitDB()
	if db == nil {
		t.Error("open database error !")
		return
	}
}
