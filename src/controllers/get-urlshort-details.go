package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gustavoteixeira8/url-shortener/src/usecases"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
)

func GetURLShortDetailsController(ctx *fiber.Ctx) error {
	req := usecases.GetURLShortDetailsRequest{
		ID: ctx.Params("id"),
	}

	useCase := usecases.NewGetURLShortDetailsUseCase()
	urlShort, err := useCase.Exec(req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.InternalError(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).JSON(utils.Ok("", urlShort))
}