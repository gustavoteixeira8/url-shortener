package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gustavoteixeira8/url-shortener/src/usecases"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
	"github.com/sirupsen/logrus"
)

func RedirectURLShortController(ctx *fiber.Ctx) error {
	req := usecases.RedirectURLShortRequest{
		Name: ctx.Params("name"),
	}

	useCase := usecases.NewRedirectURLShortUseCases()
	urlShort, err := useCase.Exec(req)

	errFormatted := utils.GetErrorInHttpFormat(err)
	if errFormatted != nil {
		return ctx.Status(errFormatted.Status).JSON(errFormatted)
	}

	err = ctx.Redirect(urlShort.URL, http.StatusMovedPermanently)
	if err != nil {
		logrus.Errorf("Error redirecting to URL %s (%v)", urlShort.URL, err)
		return err
	}

	return nil
}
