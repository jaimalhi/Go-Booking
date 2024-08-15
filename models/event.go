package models

import (
	"time"

	db "example.com/rest-api/database"
)

// `binding:"required"` forces the field to be required
type Event struct {
	ID          int64
	Name        string `binding:"required"`	
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID 	    int64
}

func (e *Event) Save() error{
	// use ? to avoid SQL injection
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
			VALUES (?, ?, ?, ?, ?)`
	// Note: Prepare() is preferred over Exec() or Query() if the statement is to be executed multiple times
	// in this scenario, we close the statement immediately so it did not make a difference, we use Exec()
	// when manipulating data in db, and Query() when fetching data
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close() // close the statement when the function returns

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err 
}

func GetAllEvents() ([]Event, error){
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	// iterate over each row using rows.Next() which returns a bool
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error){
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Event) Update() error { 
	query := `UPDATE events 
			  SET name = ?, description = ?, location = ?, date_time = ? 
			  WHERE id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err 
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID)
	return err
}