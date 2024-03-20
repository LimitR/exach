package main

import (
	page_admin "exach/internal/handlers/pages/admin"
	page_smiles "exach/internal/handlers/pages/smiles"
	page_threads "exach/internal/handlers/pages/threads"
	handler_threads "exach/internal/handlers/threads"
	service_threads "exach/internal/services/threads"
	"exach/internal/services/upload"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func main() {

	secret, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		secret["DB_HOST"], secret["DB_PORT"], secret["POSTGRES_USER"], secret["POSTGRES_PASSWORD"], secret["POSTGRES_DB"])

	db, err := sqlx.Connect(secret["DB_DRIVER"], dsn)
	if err != nil {
		log.Fatalln(err)
	}

	engine := mustache.New("./templates", ".mustache")

	engine.Reload(true)

	engine.Layout("embed")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "templates/main",
	})

	app.Static("/public", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"imageLogo":       "/public/img/logo.png",
			"imagePravdorubs": "/public/img/pravdorubs.png",
			"imageExachan":    "/public/img/exachan.png",
		}, "main")
	})

	uploader := new(upload.UploaderMultipartForm)

	threadService := service_threads.NewThreadService(db)

	threadHandler := handler_threads.NewThreadHandler(threadService, uploader)

	pageThreadHandler := page_threads.NewPageThreadHandler(threadService)

	app.Get("/admin", page_admin.LoginAdmin)

	app.Get("/new/thread", pageThreadHandler.NewThread)

	app.Get("/new/smiles/:id", page_smiles.Smiles)
	app.Get("/thread/:id", pageThreadHandler.Thread)

	// API
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Run("up", db.DB, "migrations"); err != nil {
		panic(err)
	}

	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})
	api.Post("/new/thread", threadHandler.CreateNewThread)

	app.Use(NotFound)
	log.Fatal(app.Listen(":3000"))
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("pages/404", fiber.Map{
		"logo":            "/public/img/logo.png",
		"errorText":       "Возможно тут была страница, но нам она больше не нравится",
		"imagePravdorubs": "/public/img/pravdorubs.png",
		"imageExachan":    "/public/img/exachan.png",
	}, "main")
}
