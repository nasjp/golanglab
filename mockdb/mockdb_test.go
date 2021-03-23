package mockdb_test

import (
	"database/sql"
	"testing"

	"github.com/cockroachdb/copyist"
	_ "github.com/lib/pq" // here
	"github.com/nasjp/golanglab/mockdb"
)

const dataSourceName = "postgresql://root@localhost:26888?sslmode=disable"
const resetScript = `
DROP TABLE IF EXISTS customers;
CREATE TABLE customers (id INT PRIMARY KEY, name TEXT);
INSERT INTO customers VALUES (1, 'Andy'), (2, 'Jay'), (3, 'Darin');
DROP TABLE IF EXISTS datatypes;
`

func resetDB() {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec(resetScript); err != nil {
		panic(err)
	}
}

func TestQueryName(t *testing.T) {
	copyist.Register("postgres", resetDB)
	defer copyist.Open().Close()

	db, _ := sql.Open("copyist_postgres", "postgresql://root@localhost")
	defer db.Close()

	name := mockdb.QueryName(db)
	if name != "Andy" {
		t.Error("failed test")
	}
}
