package dao

import (
	"github.com/labstack/gommon/log"
	"tickets-backend/internal/api"
	"time"
)

func CreateSchema() {
	db := GetDB().Unsafe()

	var schema = `
	CREATE TABLE tickets (
		ticket_id BIGSERIAL PRIMARY KEY,
		title TEXT,
		content TEXT,
		status  VARCHAR(255),
		author  VARCHAR(255),
		created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		modified_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	  );
	`
	db.MustExec(schema)
}

func GetTickets() ([]api.Ticket, error) {

	//CreateSchema()

	db := GetDB().Unsafe()

	stmt, err := db.PrepareNamed("SELECT * FROM tickets")

	tickets := []api.Ticket{}
	err = stmt.Select(&tickets, map[string]interface{}{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return tickets, nil
}


func PersistTicket(ticket *api.Ticket) (insert api.Ticket, err error) {

	db := GetDB().Unsafe()

	// Insert ticket and get id

	stmt, err := db.PrepareNamed("INSERT INTO tickets (title, content, status, author) VALUES (:title, :content, :status, :author) RETURNING ticket_id")

	var id int
	err = stmt.Get(&id, ticket)

	if err != nil {
		return insert, err
	}

	// Select new ticket with id in db

	err = db.Get(&insert, "SELECT * FROM tickets WHERE ticket_id=$1", id)

	if err != nil {
		return insert, err
	}

	return insert, nil
}

func UpdateTicketStatus(id int64, status string) (updated api.Ticket, err error) {

	db := GetDB().Unsafe()

	res, err := db.Exec("UPDATE tickets SET status=$1, modified_at=$2 WHERE ticket_id=$3", status, time.Now(), id)

	if err != nil {
		log.Error(err)
		return updated, err
	}

	count, err := res.RowsAffected()

	if err != nil || count != 1 {
		log.Error(err)
		return updated, err
	}

	err = db.Get(&updated, "SELECT * FROM tickets WHERE ticket_id=$1", id)

	if err != nil {
		return updated, err
	}

	return updated, nil
}

// func DeteteTicket(id) (error) {

// 	db := GetDB().Unsafe()

// 	err := db.MustExec("DELETE FROM tickets WHERE id = :id", id)

// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }