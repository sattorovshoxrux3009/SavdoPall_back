package v1

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/server/models"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func saveImage(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	const maxFileSize = 5 * 1024 * 1024 // 10MB
	var allowedExtensions = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if file.Size > maxFileSize {
		return "", fmt.Errorf("file size too large, maximum allowed size is 5MB")
	}
	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[fileExtension] {
		return "", fmt.Errorf("invalid file type, only JPG, JPEG, PNG, and WEBP are allowed")
	}
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), fileExtension)
	dst := filepath.Join("uploads", newFileName)
	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return "", err
	}
	if err := c.SaveFile(file, dst); err != nil {
		return "", err
	}
	return "/uploads/" + newFileName, nil
}

func (h *handlerV1) CreateProduct(c *fiber.Ctx) error {
	var req models.CreateProduct

	// String qiymatlar
	req.Name = c.FormValue("name")
	req.Description = c.FormValue("description")

	if len(req.Name) < 3 || len(req.Name) > 255 {
		return c.Status(400).JSON(fiber.Map{"error": "Name must be between 3 and 255 characters"})
	}
	if len(req.Description) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Description must be between 3 and 255 characters"})
	}

	// Float qiymatlar
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil || price <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or missing price"})
	}
	req.Price = price

	height, err := strconv.ParseFloat(c.FormValue("height"), 64)
	if err != nil || height <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or missing height"})
	}
	req.Height = height

	width, err := strconv.ParseFloat(c.FormValue("width"), 64)
	if err != nil || width <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or missing width"})
	}
	req.Width = width

	depth, err := strconv.ParseFloat(c.FormValue("depth"), 64)
	if err != nil || depth <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or missing depth"})
	}
	req.Depth = depth

	// Int qiymatlar
	quantity, err := strconv.Atoi(c.FormValue("quantity"))
	if err != nil || quantity < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid or missing quantity"})
	}
	req.Quantity = quantity

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	file, err := c.FormFile("image")
	if err != nil || file == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Image file is required"})
	}

	imageURL, err := saveImage(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving image"})
	}
	req.ImgUrl = imageURL

	newProduct, err := h.strg.Product().Create(c.Context(), &repo.Product{
		ImgUrl:      req.ImgUrl,
		Name:        req.Name,
		Price:       req.Price,
		Height:      req.Height,
		Width:       req.Width,
		Depth:       req.Depth,
		Quantity:    req.Quantity,
		Description: req.Description,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Product": newProduct})
}

func (h *handlerV1) GetProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	if productId != "" {
		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}
		_ = id
		menu, err := h.strg.Product().GetById(c.Context(), id)
		if err != nil || menu == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.JSON(menu)
	}
	products, err := h.strg.Product().Get(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get products"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Products": products,
	})
}
