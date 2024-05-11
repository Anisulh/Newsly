package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (h *Handler) GetHomePage(c *fiber.Ctx) error {

	handler := adaptor.HTTPHandler(templ.Handler(templates.Home()))

	return handler(c)

}
