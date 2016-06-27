package authen

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB(sqlPath string) (db *DB, err error) {
	var sqlite *sql.DB
	sqlite, err = sql.Open("sqlite3", sqlPath)
	return &DB{
		db: sqlite,
	}, err
}

func (d *DB) CreateTable() error {
	sqlStmt := `
	create table if not exists authen (
		id integer primary key,
		username text,
		password text
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertAuthen(username string, password string) error {
	sqlStmt := `
	insert into authen(username, password) values('%s', '%s')
	`
	sqlStmt = fmt.Sprintf(sqlStmt, username, password)
	_, err := d.db.Exec(sqlStmt)
	return err
}

func (d *DB) QueryAuthen(username string) (authen *Authen, err error) {
	sqlStmt := `
	select username, password from authen a where username = '?'
	`

	stmt, err := d.db.Prepare(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rUsername, password string
	err = stmt.QueryRow(username).Scan(&rUsername, &password)
	if err != nil {
		return nil, err
	}

	return &Authen{
		Username: rUsername,
		Password: password,
	}, nil
}

func (d *DB) ListAllAuthen() (authens []*Authen, err error) {
	sqlStmt := `
	select username, password from authen
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var username, password string
		rows.Scan(&username, &password)

		authens = append(authens, &Authen{
			Username: username,
			Password: password,
		})
	}

	return authens, nil
}

func (d *DB) Close() {
	d.db.Close()
}
