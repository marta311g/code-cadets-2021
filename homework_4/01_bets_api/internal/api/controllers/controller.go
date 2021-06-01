package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller implements handlers for web server requests.
type Controller struct {
	dbService  DbService
	dbExecutor sqlite.DatabaseExecutor
}

// NewController creates a new instance of Controller
func NewController(dbService DbService, dbExecutor sqlite.DatabaseExecutor) *Controller {
	return &Controller{
		dbService:  dbService,
		dbExecutor: dbExecutor,
	}
}

// GetBetByID handlers gets the bet with id.
func (e *Controller) GetBetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// get bet from db??????????????????????
		// get id from url
		// query from db by id

		// marshal it

		// send it
		// send status code

		id := ctx.Param("id")

		betResponse, err := e.dbService.GetBetByID(id, e.dbExecutor) //get id from uri
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
	}
}

func (e *Controller) GetBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
