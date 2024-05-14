package controller

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/oarielg/TarotGo/cards"
	"github.com/oarielg/TarotGo/view"
)

func IndexHandler(c echo.Context) error {
	component := view.Index()
	return component.Render(context.Background(), c.Response().Writer)
}

func ReadHandler(c echo.Context) error {
	deck := cards.NewDeck()
	deck = cards.SuffleDeck(deck)

	hand := cards.DrawCards(deck, 3)

	component := view.Read(hand)
	return component.Render(context.Background(), c.Response().Writer)
}
