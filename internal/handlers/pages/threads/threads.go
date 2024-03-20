package page_threads

import "github.com/gofiber/fiber/v2"

func NewThread(c *fiber.Ctx) error {
	apiNewThread := "/api/new/thread"

	return c.Render("pages/new_post", fiber.Map{
		"apiNewThread":    apiNewThread,
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",
	}, "main")
}
