package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
)

func newApp() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world"})
	})
	return app
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := newApp()
	app.Hooks().OnPostShutdown(func(err error) error {
		if err != nil {
			log.Printf("Graceful shutdown gagal: %v", err)
		} else {
			log.Print("Graceful shutdown berhasil")
		}
		return nil
	})

	if err := app.Listen(":3000", fiber.ListenConfig{
		GracefulContext: ctx,
		ShutdownTimeout: 10 * time.Second,
	}); err != nil {
		log.Fatal(err)
	}
}
