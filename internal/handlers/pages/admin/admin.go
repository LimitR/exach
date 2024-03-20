package page_admin

import "github.com/gofiber/fiber/v2"

func LoginAdmin(c *fiber.Ctx) error {
	return c.Render("pages/admin", fiber.Map{
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",
	}, "main")
}
