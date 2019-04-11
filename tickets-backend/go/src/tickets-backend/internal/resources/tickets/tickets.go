package tickets

import (
	"github.com/labstack/echo"
	"net/http"
	"tickets-backend/internal/dao"
	"tickets-backend/internal/api"
)

/**
 * Declaration of routes
 */
func Route(group *echo.Group) {
	group.GET("/", GetTickets)
	group.POST("/", InsertTicket)
	group.POST("/update", UpdateTicketStatus)
	group.POST("/delete", DeleteTicket)
}

/**
 * Route to list all tickets
 */
func GetTickets(c echo.Context) error {
	tickets, err :=  dao.GetTickets()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, tickets)
}

/**
 * Route to insert ticket
 */
func InsertTicket(c echo.Context) (err error) {

	ticket := new(api.Ticket)
	if err = c.Bind(ticket); err != nil {
		return
	}

	inserted, err := dao.PersistTicket(ticket)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, inserted)
}

/**
 * Route to update ticket status
 */
func UpdateTicketStatus(c echo.Context) (err error) {

	ticket := new(api.Ticket)
	if err = c.Bind(ticket); err != nil {
		return
	}

	updated, err := dao.UpdateTicketStatus(ticket.TicketId, ticket.Status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, updated)
}

/**
 * Route to delete ticket
 */
func DeleteTicket(c echo.Context) (err error) {

	ticket := new(api.Ticket)
	if err = c.Bind(ticket); err != nil {
		return
	}

	err = dao.DeleteTicket(ticket.TicketId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "deleted")
}