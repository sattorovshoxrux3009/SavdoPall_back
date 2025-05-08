package server

import (
	"time"

	v1 "GitHub.com/sattorovshoxrux3009/SavdoPall_back/server/v1"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Options struct {
	Strg storage.StorageI
}

func NewServer(opts *Options) *fiber.App {
	// app := fiber.New(fiber.Config{
	// 	// HTTPS uchun TLS konfiguratsiyasi
	// 	DisableKeepalive: false,
	// })
	app := fiber.New()
	// IP log middleware
	app.Use(func(c *fiber.Ctx) error {
		clientIP := c.IP()
		requestTime := time.Now().Format("2006-01-02 15:04:05") // Yil-oy-kun soat:minut:sekund
		println("Yangi so‘rov! IP:", clientIP, "Vaqt:", requestTime)
		return c.Next()
	})

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Authorization",
		AllowOriginsFunc: func(origin string) bool { return true }, // OPTIONS muammosini hal qiladi
		// AllowCredentials: true,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        60,              // Maksimal 100 ta so‘rov
		Expiration: 1 * time.Minute, // 1 daqiqa ichida hisoblash
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Foydalanuvchi IP manzili bo‘yicha cheklash
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "So‘rovlar soni cheklangan, keyinroq urinib ko‘ring.",
			})
		},
	}))

	// Handler
	handler := v1.New(&v1.HandlerV1{
		Strg: opts.Strg,
	})
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(204) // No Content
	})
	app.Static("/uploads", "./uploads")

	app.Post("/v1/admin", handler.CreateAdmin)
	app.Post("/v1/login", handler.Login)

	app.Post("/v1/product", handler.AuthMiddleware(), handler.CreateProduct)
	app.Patch("/v1/product/:id", handler.AuthMiddleware(), handler.UpdateProduct)
	app.Delete("/v1/product/:id", handler.AuthMiddleware(), handler.DeleteProduct)

	app.Get("/v1/product", handler.GetProduct)
	app.Get("/v1/product/:id", handler.GetProduct)

	return app
}
