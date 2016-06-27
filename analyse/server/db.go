package analyse

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
	create table if not exists event (
		id integer primary key,
		target_id integer,
		event text,
		time integer
	)	
	`

	_, err := d.db.Exec(sqlStmt)

	return err
}

func (d *DB) InsertEvent(targetID int, event string, time int64) error {
	sqlStmt := `
	insert into event(target_id, event, time) values(%d, '%s', %d)
	`
	_, err := d.db.Exec(fmt.Sprintf(sqlStmt, targetID, event, time))
	return err
}

func (d *DB) ListAllEvent() (events []*Event, err error) {
	sqlStmt := `
	select target_id, event, time from event
	`

	rows, err := d.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var targetID int
		var event string
		var time int64

		rows.Scan(&targetID, &event, &time)

		events = append(events, &Event{
			TargetID: targetID,
			Event:    event,
			Time:     time,
		})
	}

	return events, nil
}

func (d *DB) Close() {
	d.db.Close()
}
