package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gustavoteixeira8/url-shortener/src/usecases"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
)

func GetURLShortDetailsController(ctx *fiber.Ctx) error {
	req := usecases.GetURLShortDetailsRequest{
		Name: ctx.Params("name"),
	}

	useCase := usecases.NewGetURLShortDetailsUseCase()
	urlShort, err := useCase.Exec(req)

	errFormatted := utils.GetErrorInHttpFormat(err)
	if errFormatted != nil {
		return ctx.Status(errFormatted.Status).JSON(errFormatted)
	}

	return ctx.Status(http.StatusOK).JSON(utils.Ok("", urlShort))
}
