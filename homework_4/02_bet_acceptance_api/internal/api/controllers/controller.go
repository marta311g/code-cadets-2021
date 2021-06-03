package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betValidator BetValidator
	betService   BetService
}

// NewController creates a new instance of Controller
func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

// InsertBet handlers insert bet request.
func (e *Controller) InsertBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betInsertRequestDto models.BetInsertRequestDto
		err := ctx.ShouldBindWith(&betInsertRequestDto, binding.JSON)
		if err != nil {
			log.Println("error")
			ctx.String(http.StatusBadRequest, "request is not valid.")
			return
		}

		if !e.betValidator.BetIsValid(betInsertRequestDto) {
			ctx.String(http.StatusBadRequest, "request is not valid.")
			return
		}

		err = e.betService.Publisher(betInsertRequestDto.CustomerId, betInsertRequestDto.SelectionId, betInsertRequestDto.SelectionCoefficient, betInsertRequestDto.Payment)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
