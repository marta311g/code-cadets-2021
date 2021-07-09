package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller implements handlers for web server requests.
type Controller struct {
	dbService  DbService
}

// NewController creates a new instance of Controller
func NewController(dbService DbService) *Controller {
	return &Controller{
		dbService:  dbService,
	}
}

// GetBetByID handlers gets the bet with id.
func (e *Controller) GetBetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		betResponse, err := e.dbService.GetBetByID(id) //get id from uri
		if err != nil {
			log.Println(err)
			ctx.String(http.StatusNotFound, "no bet for this id")
			return
		}

		ctx.JSON(http.StatusOK, betResponse)
	}
}

func (e *Controller) GetBetsByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")

		betsResponse, err := e.dbService.GetBetsByUser(userId)
		if err != nil {
			log.Println(err)
			ctx.String(http.StatusNotFound, "no bets for this user")
			return
		}

		ctx.JSON(http.StatusOK, betsResponse)
	}
}

func (e *Controller) GetBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")

		betsResponse, err := e.dbService.GetBetsByStatus(status)
		if err != nil {
			log.Println(err)
			ctx.String(http.StatusNotFound, "no bets with this status")
			return
		}

		ctx.JSON(http.StatusOK, betsResponse)
	}
}
