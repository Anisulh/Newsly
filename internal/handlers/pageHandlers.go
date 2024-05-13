package handlers

import (
	"Newsly/web/templates/pages"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)



func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func (h *Handler) GetHomePage(c *fiber.Ctx) error {
	return Render(c, pages.HomePage())
}

func (h *Handler) GetLoginPage(c *fiber.Ctx) error {
	return Render(c, pages.LoginPage())
}

func (h *Handler) GetRegisterPage(c *fiber.Ctx) error {
	return Render(c, pages.RegisterPage())
}
