package logging

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
	create table if not exists log (
		id integer primary key,
		addr text,
		namespace text,
		cause text
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertLog(addr string, namespace string, cause string) (log *Log, err error) {
	sqlStmt := `
	insert into log(addr, namespace, cause) values('%s', '%s', '%s')
	`
	sqlStmt = fmt.Sprintf(sqlStmt, addr, namespace, cause)
	result, err := d.db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &Log{
		ID:        int(id),
		Addr:      addr,
		Namespace: namespace,
		Cause:     cause,
	}, nil
}

func (d *DB) ListAllLog() (logs []*Log, err error) {
	sqlStmt := `
	select id, addr, namespace, cause from log
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var addr, namespace, cause string

		rows.Scan(&id, &addr, &namespace, &cause)

		logs = append(logs, &Log{
			ID:        id,
			Addr:      addr,
			Namespace: namespace,
			Cause:     cause,
		})
	}

	return logs, nil
}

func (d *DB) Close() {
	d.db.Close()
}
