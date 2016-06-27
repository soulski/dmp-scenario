package user

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
	create table if not exists user (
		id integer primary key,
		username text,
		email text,
		address text
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertUser(username string, email string, address string) (user *User, err error) {
	sqlStmt := `
	insert into user(username, email, address) values('%s', '%s', '%s')
	`
	sqlStmt = fmt.Sprintf(sqlStmt, username, email, address)
	result, err := d.db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &User{
		ID:       int(id),
		Username: username,
		Email:    email,
		Address:  address,
	}, nil
}

func (d *DB) ListAllUser() (users []*User, err error) {
	sqlStmt := `
	select id, username, email, address from user
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, email, addr string

		rows.Scan(&id, &username, &email, &addr)

		users = append(users, &User{
			ID:       id,
			Username: username,
			Email:    email,
			Address:  addr,
		})
	}

	return users, nil
}

func (d *DB) Close() {
	d.db.Close()
}
