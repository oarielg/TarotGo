package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oarielg/TarotGo/cards"
	"github.com/oarielg/TarotGo/view"
)

func IndexHandler(c echo.Context) error {
	component := view.Index()
	return view.Primary(component).Render(context.Background(), c.Response().Writer)
}

func ReadHandler(c echo.Context) error {
	deck := cards.NewDeck()
	deck = cards.SuffleDeck(deck)

	hand := cards.DrawCards(deck, 3)

	component := view.Read(hand)
	return component.Render(context.Background(), c.Response().Writer)
}

func CardHandler(c echo.Context) error {
	id := c.Param("id")
	card := cards.GetCard(id)
	if card.Name == "" {
		return echo.NewHTTPError(http.StatusNotFound, "This page does not exist.")
	}

	component := view.Card(cards.CardDrawn{Card: card, Reversed: false})
	return view.Primary(component).Render(context.Background(), c.Response().Writer)
}

func CardsHandler(c echo.Context) error {
	deck := cards.AllCards()

	component := view.Cards(deck)
	return view.Primary(component).Render(context.Background(), c.Response().Writer)
}
