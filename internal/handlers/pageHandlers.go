package handlers

import (
	// "Newsly/internal/utils"
	"Newsly/internal/models"
	"Newsly/internal/utils"
	"Newsly/web/templates/pages"
	"log"

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

func (h *Handler) GetInterestsPage(c *fiber.Ctx) error {
	userID := c.Locals("account").(uint) // Ensure type assertion to uint

	// Fetch user details from the database using userID
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.Redirect("/login")
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}

	var categoryMapping = map[string]utils.Category{
		"ml": {
			Title:       "Artificial Intelligence & Machine Learning",
			Description: "Focus on advancements in algorithms, neural networks, and data-driven decision-making.",
		},
		"quantum": {
			Title:       "Quantum Physics & Quantum Computing",
			Description: "Studies in quantum mechanics, quantum algorithms, and computing architectures.",
		},
		"neuroscience": {
			Title:       "Neuroscience & Cognitive Science",
			Description: "Explores brain function, neural processes, and the science of cognition.",
		},
		"genetics": {
			Title:       "Genetics & Genomics",
			Description: "Research on DNA sequencing, gene expression, and genetic engineering.",
		},
		"renewables": {
			Title:       "Renewable Energy & Sustainability",
			Description: "Innovations in solar, wind, and sustainable energy technologies with environmental impact studies.",
		},
		"biotech": {
			Title:       "Biotechnology & Medical Research",
			Description: "Advancements in drug discovery, biomedical engineering, and clinical studies.",
		},
		"astrophysics": {
			Title:       "Astrophysics & Cosmology",
			Description: "Research on the universe, celestial bodies, and cosmic phenomena.",
		},
		"robotics": {
			Title:       "Robotics & Automation",
			Description: "Design, control systems, and the integration of robotics in various applications.",
		},
		"materials": {
			Title:       "Materials Science & Nanotechnology",
			Description: "Development of new materials, nanostructures, and their practical applications.",
		},
		"data_science": {
			Title:       "Data Science & Big Data Analytics",
			Description: "Methodologies for analyzing vast datasets, data mining, and statistical modeling.",
		},
		"cs": {
			Title:       "Computer Science & Software Engineering",
			Description: "Topics including programming languages, system architectures, and development methodologies.",
		},
		"economics": {
			Title:       "Economics & Behavioral Science",
			Description: "Studies on economic theory, market behaviors, and the psychology behind decision-making.",
		},
		"environment": {
			Title:       "Environmental Science & Climate Change",
			Description: "Research on global warming, ecosystem management, and environmental policy.",
		},
		"social": {
			Title:       "Social Sciences & Humanities",
			Description: "Explorations in sociology, anthropology, political science, and cultural studies.",
		},
		"engineering": {
			Title:       "Engineering & Technological Innovation",
			Description: "Advancements in civil, mechanical, electrical, and chemical engineering.",
		},
		"blockchain": {
			Title:       "Blockchain & Cryptocurrency",
			Description: "Research on distributed ledger technologies, digital currencies, and decentralized finance.",
		},
		"medicine": {
			Title:       "Medicine & Healthcare",
			Description: "Studies in medical science, healthcare technologies, and patient care innovations.",
		},
	}

	extendedData := utils.InterestTopicsData{
		BaseData: utils.BaseData{
			IsAuth:  true,
			Account: &user,
		},
		Categories: categoryMapping,
	}
	return Render(c, pages.InterestTopics(extendedData))
}

func (h *Handler) GetFeedPage(c *fiber.Ctx) error {
	userID := c.Locals("account").(uint) // Ensure type assertion to uint


	// Fetch user details from the database using userID
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.Redirect("/login")
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	// get the user's interests
	var savedCategories []models.Category
	if err := h.DB.Model(&user).Association("CategoryInterests").Find(&savedCategories); err != nil {
		log.Println("Error fetching user interests:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching user interests")
	}

	var userInterests []string
	for _, cat := range savedCategories {
		userInterests = append(userInterests, cat.Key)
	}
	builtQuery := utils.BuildArxivQuery(userInterests)
	arxivResp, err := utils.FetchArxivPapers(builtQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching papers")
	}

	papers := arxivResp.Entries

	extendedData := utils.FeedData{
		BaseData: utils.BaseData{
			IsAuth:  true,
			Account: &user,
		},
		Papers: papers,
	}
	return Render(c, pages.FeedPage(extendedData))
}

func (h *Handler) GetSpecificNewsPage(c *fiber.Ctx) error {
	account, ok := c.Locals("Account").(*models.User)
	if !ok {
		// Handle the case where the type assertion fails
		return c.Status(fiber.StatusInternalServerError).SendString("Account information is not available")
	}
	extendedData := utils.FeedData{
		BaseData: utils.BaseData{
			IsAuth:  true,
			Account: account,
		},
		// Articles: articles,
	}
	return Render(c, pages.FeedPage(extendedData))
}
