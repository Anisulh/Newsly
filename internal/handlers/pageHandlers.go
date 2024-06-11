package handlers

import (
	// "Newsly/internal/utils"
	"Newsly/internal/models"
	"Newsly/internal/utils"
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

func (h *Handler) GetFeedPage(c *fiber.Ctx) error {
	userID := c.Locals("account").(uint) // Ensure type assertion to uint

	// Optional: Fetch user details from the database using userID
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	resp, err := utils.FetchNews(h.NewsAPIKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching news")
	}
	articles := resp.Articles
	println(&articles)
	data := utils.BaseData{
		IsAuth:   true,
		Account:  &user, // Assuming you want to use the user model
		Articles: articles,
	}
	return Render(c, pages.FeedPage(data))
}

func (h *Handler) GetSpecificNewsPage(c *fiber.Ctx) error {
	account, ok := c.Locals("Account").(*models.User)
	if !ok {
		// Handle the case where the type assertion fails
		return c.Status(fiber.StatusInternalServerError).SendString("Account information is not available")
	}
	data := utils.BaseData{
		IsAuth:  true,
		Account: account,
	}
	return Render(c, pages.FeedPage(data))
}
