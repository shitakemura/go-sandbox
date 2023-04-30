package testutil

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	db, err := sql.Open(
		"mysql",
		"todo:todo@tcp(127.0.0.1:33306)/todo?parseTime=true",
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	return sqlx.NewDb(db, "mysql")
}
