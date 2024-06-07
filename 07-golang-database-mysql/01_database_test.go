package golang_database_mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
}

func TestOpenConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}
