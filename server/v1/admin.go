package v1

import (
	"context"
	"time"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/server/models"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("shoxrux2004$")

func generateJWT(adminID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"username": username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}

func (h *handlerV1) CreateAdmin(c *fiber.Ctx) error {
	// Request body'dan CreateAdmin struct'ini olish
	var req models.CreateAdmin

	// Validatsiya qilish
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	oldAdmin, err := h.strg.Admin().GetByUName(c.Context(), req.Username)
	if err == nil && oldAdmin != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "This username already exists",
		})
	}

	// Adminni yaratish uchun yangi Admin modeli
	admin := repo.Admin{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Username:     req.Username,
		PasswordHash: hashPassword(req.Password),
	}

	// Adminni bazaga saqlash
	if err := h.strg.Admin().Create(c.Context(), &admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create admin",
		})
	}

	// Yaratilgan adminni javob sifatida yuborish
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Admin created successfully",
	})
}

func (h *handlerV1) Login(c *fiber.Ctx) error {
	var req models.Login
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	admin, err := h.strg.Admin().GetByUName(context.TODO(), req.Username)
	if err != nil || admin == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username or password error",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username or password error",
		})
	}
	token, err := generateJWT(admin.Id, admin.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	updates := make(map[string]interface{})
	updates["token"] = token
	err = h.strg.Admin().Update(c.Context(), int(admin.Id), updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
