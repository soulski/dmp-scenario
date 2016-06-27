package search

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
	create table if not exists record (
		id integer primary key,
		id_ref integer,
		token text
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertRecord(idRef int, token string) error {
	sqlStmt := `
	insert into record(id_ref, token) values(%d, '%s')
	`
	_, err := d.db.Exec(fmt.Sprintf(sqlStmt, idRef, token))
	return err
}

func (d *DB) ListAllRecord() (records []*Record, err error) {
	sqlStmt := `
	select id, id_ref, token from record
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, id_ref int
		var token string

		rows.Scan(&id, &id_ref, &token)

		records = append(records, &Record{
			ID:    id,
			IDRef: id_ref,
			Token: token,
		})
	}

	return records, nil
}

func (d *DB) Close() {
	d.db.Close()
}
