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

func (p *PageThreadHandler) GetThreads(c *fiber.Ctx) error {
	board := c.Params("board", "")
	threads := p.serviceThread.GetThreadsFromBoard(board, 100)

	html := make([]interface{}, 0, 10)
	if len(threads) != 0 {
		for _, thread := range threads {
			html = append(html, fiber.Map{
				"text":      textmatching.MatchTextToHTML(thread.Text),
				"id":        thread.Id,
				"img":       thread.Img,
				"boardName": thread.Board,
				"postTitle": thread.Head,
			})
		}
	}

	return c.Render("pages/board", fiber.Map{
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",

		"threads": html,

		"board": `<li><a href="/b">Бред</a></li>`,
	}, "main")
}

func (p *PageThreadHandler) Thread(c *fiber.Ctx) error {
	id := c.Params("id")
	board := c.Params("board", "")

	idN, _ := strconv.Atoi(id)

	thread := p.serviceThread.GetThreadById(idN, board)

	thread.Text = textmatching.MatchTextToHTML(thread.Text)

	result := p.serviceThread.GetCommentsByThreadId(thread.Id)

	html := make([]interface{}, 0, 10)

	if len(result) != 0 {
		for _, thread := range result {
			html = append(html, fiber.Map{
				"text": textmatching.MatchTextToHTML(thread.Text),
				"id":   thread.Id,
			})
		}
	}

	return c.Render("pages/thread", fiber.Map{
		"logo":            "/public/img/logo.png",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",

		"postId":    id,
		"postTitle": thread.Head,
		"postText":  thread.Text,
		"postImg":   textmatching.LinkToImage(thread.Img),

		"comments": html,

		"board": `<li><a href="/b">Бред</a></li>`,
	}, "main")
}
