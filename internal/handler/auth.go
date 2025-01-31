package handler

import (
	"mystats-server/internal/auth"
	"mystats-server/internal/ent"
	"mystats-server/internal/ent/user"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*"`
	Name     string `json:"name" validate:"required,min=2"`
	Team     string `json:"team,omitempty" validate:"omitempty,min=2"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	user, err := h.client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Only(c.Context())
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if !auth.CheckPassword(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	var req SignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request format",
			"details": "Request body must be valid JSON",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": "Password must be at least 8 characters and contain special characters",
		})
	}

	exists, err := h.client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Exist(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error",
			"details": "Failed to check email availability",
		})
	}
	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   "Email already registered",
			"details": "Please use a different email or try logging in",
		})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user, err := h.client.User.Create().
		SetEmail(req.Email).
		SetPassword(hashedPassword).
		SetName(req.Name).
		SetTeam(req.Team).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"token":   token,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"team":  user.Team,
		},
	})
}

func (h *Handler) CreateTestUser(c *fiber.Ctx) error {
	hashedPassword, err := auth.HashPassword("test1234")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user, err := h.client.User.Create().
		SetEmail("test@example.com").
		SetPassword(hashedPassword).
		SetName("Test User").
		SetTeam("Test Team").
		Save(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create test user",
		})
	}

	return c.JSON(user)
}

func (h *Handler) ValidateToken(c *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	// Remove "Bearer " prefix
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}

	// Validate token
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Get user from database
	user, err := h.client.User.Get(c.Context(), claims.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	return c.JSON(fiber.Map{
		"valid": true,
		"user":  user,
	})
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.client.User.Get(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user profile",
		})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"team":  user.Team,
		},
	})
}
