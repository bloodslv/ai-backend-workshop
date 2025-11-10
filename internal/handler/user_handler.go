package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"kbtg.tech/ai-backend-workshop/internal/domain"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userUseCase domain.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"count": len(users),
	})
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req domain.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userUseCase.CreateUser(req)
	if err != nil {
		if err.Error() == "first name, last name, and email are required" ||
			err.Error() == "user with this email already exists" {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": user,
	})
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req domain.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userUseCase.UpdateUser(uint(id), req)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		if err.Error() == "user with this email already exists" {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = h.userUseCase.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
