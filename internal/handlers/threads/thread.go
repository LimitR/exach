package handler_threads

import (
	"exach/internal/database/models"
	service_threads "exach/internal/services/threads"
	"exach/internal/services/upload"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ThreadHandler struct {
	threadService *service_threads.ThreadService
	uploader      *upload.UploaderMultipartForm
}

func NewThreadHandler(threadService *service_threads.ThreadService, uploader *upload.UploaderMultipartForm) *ThreadHandler {
	return &ThreadHandler{
		threadService: threadService,
		uploader:      uploader,
	}
}

func (t *ThreadHandler) CreateNewThread(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {

		thread := new(models.Thread)

		if text := form.Value["text"]; len(text) == 0 {
			return c.Status(400).Redirect("/notvalid")
		} else {
			thread.Text = text[0]
		}
		if title := form.Value["title"]; len(title) != 0 {
			thread.Head = title[0]
		}
		if email := form.Value["email"]; len(email) != 0 {
			thread.UserName = email[0]
		}
		if file := form.File["img"]; len(file) > 0 {

			if file[0] != nil && file[0].Filename != "" {
				thread.Img, err = t.uploader.SaveFileToDisk(*file[0])
				if err != nil {
					return c.Status(400).Redirect("/notvalid")
				}

			}
		}
		if threadId := form.Value["id"]; len(threadId) != 0 {
			thread.ThreadId = threadId[0]
		}
		id, err := t.threadService.CreateThread(*thread)
		if err != nil {
			fmt.Println(err)
		}
		if thread.UserName != "" {
			return c.Status(200).Redirect("/thread/" + strconv.Itoa(id))
		} else {
			c.Status(200).Redirect("/thread/" + strconv.Itoa(id))
		}
	}

	return c.Status(400).Redirect("/notvalid")
}

func (t *ThreadHandler) GetNewComments(c *fiber.Ctx) error {
	threadId := c.Query("threadId")

	if threadId == "" {
		return c.Status(400).Redirect("/notvalid")
	}

	result := t.threadService.GetCommentsByThreadId(0)

	if len(result) == 0 {
		return c.Status(200).SendString("")
	}

	html := ""

	for _, thread := range result {
		html += fmt.Sprintf(`<div class="comment" id="c%d">
	<p>%s</p>
    <span title="Показать" class="show_image" onclick="show_hide(this)"></span>
    <span class="info">Комментарий №%d <a href="#" onclick="reply(this); document.getElementById('f1').style.display = 'block';">ответить</a></span>
	</div>`, thread.Id, thread.Text, thread.Id)
	}

	return c.Status(200).SendString(html)
}
