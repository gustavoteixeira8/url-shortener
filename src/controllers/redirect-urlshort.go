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
		ID: ctx.Params("id"),
	}

	useCase := usecases.NewRedirectURLShortUseCases()
	urlShort, err := useCase.Exec(req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.InternalError(err.Error(), nil))
	}

	err = ctx.Redirect(urlShort.URL, http.StatusMovedPermanently)
	if err != nil {
		logrus.Errorf("Error redirecting to URL %s (%v)", urlShort.URL, err)
		return err
	}

	return nil
}
