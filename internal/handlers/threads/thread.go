package handler_threads

import (
	"exach/internal/database/models"
	service_threads "exach/internal/services/threads"
	"exach/internal/services/upload"

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
		if title := form.Value["title"]; len(title) == 0 {
			return c.Status(400).Redirect("/notvalid")
		} else {
			thread.Head = title[0]
		}
		if email := form.Value["email"]; len(email) == 0 {
			return c.Status(400).Redirect("/notvalid")
		} else {
			thread.Password_hash = email[0]
		}
		if file := form.File["img"]; len(file) > 0 {

			if file[0] != nil && file[0].Filename != "" {
				pathFile, err := t.uploader.SaveFileToDisk(*file[0])
				if err != nil {
					return c.Status(400).Redirect("/notvalid")
				}

				thread.Img = pathFile

				t.threadService.CreateThread(*thread)

				return c.Status(200).Redirect("/")
			}
		}
	}

	return c.Status(400).Redirect("/notvalid")
}
