package store

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
	create table if not exists item (
		id integer primary key,
		name text,
		description text,
		price integer
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertItem(name string, description string, price int) (item *Item, err error) {
	sqlStmt := `
	insert into item(name, description, price) values('%s', '%s', %d)
	`
	sqlStmt = fmt.Sprintf(sqlStmt, name, description, price)
	result, err := d.db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &Item{
		ID:          int(id),
		Name:        name,
		Description: description,
		Price:       price,
	}, nil
}

func (d *DB) ListAllItem() (items []*Item, err error) {
	sqlStmt := `
	select id, name, description, price from item
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, price int
		var name, description string

		rows.Scan(&id, &name, &description, &price)

		items = append(items, &Item{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
		})
	}

	return items, nil
}

func (d *DB) Close() {
	d.db.Close()
}
