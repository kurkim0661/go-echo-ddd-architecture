package command

import (
	"net/http"
	"strconv"

	"myapp/api/application"
	"myapp/api/domain/dto"
	"myapp/api/interfaces/client/command/response"

	"github.com/labstack/echo/v4"
)

type CardController struct {
	application.CardService
}

func (CardController) NewCardController(service application.CardService) *CardController {
	return &CardController{service}
}

func (cardController *CardController) Init(e *echo.Group) {
	e.POST("", cardController.CreateCard)
	e.DELETE("", cardController.DeleteCard)
}

func (cardController *CardController) DeleteCard(c echo.Context) error {
	cardId, err := strconv.Atoi(c.QueryParam("cardId"))
	if err != nil {
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, err)
	}

	userId, err := strconv.Atoi(c.QueryParam("userId"))
	if err != nil {
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, err)
	}

	cardController.CardService.DeleteCard(cardId, userId)

	return response.ReturnApiSuccess(c, http.StatusNoContent, nil)
}

func (cardController *CardController) CreateCard(c echo.Context) error {
	cardDto := &dto.CardDto{}
	err := c.Bind(cardDto)
	if err != nil {
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, err)
	}
	card, err := cardController.CardService.CreateCard(cardDto)
	if err != nil {
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	return response.ReturnApiSuccess(c, http.StatusCreated, card)
}
