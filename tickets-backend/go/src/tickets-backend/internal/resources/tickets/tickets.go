package tickets

import (
	"github.com/labstack/echo"
	"net/http"
	"tickets-backend/internal/dao"
	"tickets-backend/internal/api"
	// "encoding/json"
	// "io/ioutil"
)

func Route(group *echo.Group) {
	group.GET("/", GetTickets)
	group.POST("/", InsertTicket)
	group.POST("/update", UpdateTicketStatus)
	// group.DELETE("/:id", DeleteTicket)
}


func GetTickets(c echo.Context) error {
	tickets, err :=  dao.GetTickets()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, tickets)
}


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

// func DeleteTicket(c echo.Context) error {
// 	err := dao.DeleteTicket(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	return c.JSON(http.StatusOK, ticket)
// }