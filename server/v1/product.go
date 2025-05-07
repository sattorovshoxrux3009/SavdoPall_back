package v1

import (
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/server/models"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"github.com/gofiber/fiber/v2"
)

func (h *handlerV1) CreateProduct(ctx *fiber.Ctx) error {
	var req models.CreateProduct
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newProduct, err := h.strg.Product().Create(ctx.Context(), &repo.Product{
		ImgUrl:      req.ImgUrl,
		Name:        req.Name,
		Price:       req.Price,
		Height:      req.Height,
		Width:       req.Width,
		Depth:       req.Depth,
		Quantity:    req.Quantity,
		Left:        req.Left,
		Description: req.Description,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"Product": newProduct})
}
