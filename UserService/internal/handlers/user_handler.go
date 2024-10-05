package handlers

import (
	"UserService/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	// Lấy tham số id từ URL dưới dạng chuỗi
	idParam := c.Params("id")

	// Chuyển đổi từ chuỗi sang uint
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user, err := h.UserService.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	c.Locals("data", user)
	return nil
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	user, err := h.UserService.CreateUser(req.Name, req.Age)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Locals("data", user)
	c.Status(fiber.StatusCreated)
	return nil
}
