package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nasjp/golanglab/sqlboiler/datamodels"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const connectionTemplate = "%s:%s@(%s:%s)/%s?parseTime=true&tls=%t&multiStatements=true"

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		connectionTemplate,
		"root",
		"",
		"db",
		"3306",
		"app",
		false,
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	boil.DebugMode = true
	u := &datamodels.User{
		CompanyID: 1,
		Name:      "bob",
	}

	db, err := connect()
	if err != nil {
		return err
	}
	if err := u.Insert(db, boil.Infer()); err != nil {
		return err
	}

	us, err := datamodels.Users().All(db)
	if err != nil {
		return err
	}

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}

	us, err = datamodels.Users(qm.Where("ID = ?", 1)).All(db)

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}

	us, err = datamodels.Users(qm.Limit(3)).All(db)

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}

	us, err = datamodels.Users(qm.Where("deleted_at = ?", time.Time{})).All(db)

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}

	// if _, err := datamodels.Users().DeleteAll(db); err != nil {
	// 	return err
	// }
	return nil
}
