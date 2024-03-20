package page_threads

import (
	service_threads "exach/internal/services/threads"
	textmatching "exach/pkg/textMatching"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PageThreadHandler struct {
	serviceThread *service_threads.ThreadService
}

func NewPageThreadHandler(serviceThread *service_threads.ThreadService) *PageThreadHandler {
	return &PageThreadHandler{
		serviceThread: serviceThread,
	}
}

func (p *PageThreadHandler) NewThread(c *fiber.Ctx) error {
	apiNewThread := "/api/new/thread"

	return c.Render("pages/new_post", fiber.Map{
		"apiNewThread":    apiNewThread,
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",
	}, "main")
}

func (p *PageThreadHandler) Thread(c *fiber.Ctx) error {
	id := c.Params("id")

	idN, _ := strconv.Atoi(id)

	thread := p.serviceThread.GetThreadById(idN)

	thread.Text = textmatching.MatchTextToHTML(thread.Text)

	return c.Render("pages/thread", fiber.Map{
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",

		"postId":    id,
		"postTitle": thread.Head,
		"postText":  thread.Text,
		"postImg":   textmatching.LinkToImage(thread.Img),
	}, "main")
}
