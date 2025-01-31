package handler

import (
	"time"

	"mystats-server/internal/ent/batterrecord"
	"mystats-server/internal/ent/game"
	"mystats-server/internal/ent/pitcherrecord"
	"mystats-server/internal/ent/user"

	"github.com/gofiber/fiber/v2"
)

type GetRecordsRequest struct {
	Year  int    `query:"year" validate:"required"`
	Month int    `query:"month" validate:"required,min=1,max=12"`
	Type  string `query:"type" validate:"required,oneof=pitcher batter"`
}

func (h *Handler) GetRecords(c *fiber.Ctx) error {
	// 쿼리 파라미터 파싱
	var req GetRecordsRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	// 유효성 검사
	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	// 사용자 ID 가져오기
	userID := c.Locals("userID").(int)

	// 시작일과 종료일 계산
	startDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	if req.Type == "pitcher" {
		// 투수 기록 조회
		records, err := h.client.PitcherRecord.Query().
			Where(
				pitcherrecord.HasUserWith(user.ID(userID)),
				pitcherrecord.HasGameWith(
					game.DateGTE(startDate),
					game.DateLT(endDate),
				),
			).
			WithGame().
			All(c.Context())
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch pitcher records",
			})
		}

		return c.JSON(fiber.Map{
			"records": records,
		})

	} else {
		// 타자 기록 조회
		records, err := h.client.BatterRecord.Query().
			Where(
				batterrecord.HasUserWith(user.ID(userID)),
				batterrecord.HasGameWith(
					game.DateGTE(startDate),
					game.DateLT(endDate),
				),
			).
			WithGame().
			All(c.Context())
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch batter records",
			})
		}

		return c.JSON(fiber.Map{
			"records": records,
		})
	}
}

func (h *Handler) CreateTestGame(c *fiber.Ctx) error {
	// 테스트 유저 가져오기
	user, err := h.client.User.Query().
		Where(user.Email("test@example.com")).
		Only(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Test user not found",
		})
	}

	// 테스트 게임 생성
	game, err := h.client.Game.Create().
		SetDate(time.Now()).
		SetOpponent("테스트팀").
		SetLocation("테스트구장").
		SetResult("승").
		SetMyScore(5).
		SetOpponentScore(2).
		SetUser(user).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create test game",
		})
	}

	// 투수 기록 생성
	_, err = h.client.PitcherRecord.Create().
		SetGames(1).
		SetEra(2.25).
		SetWins(1).
		SetInnings(8).
		SetStrikeouts(7).
		SetHitsAllowed(5).
		SetEarnedRuns(2).
		SetUser(user).
		SetGame(game).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create pitcher record",
		})
	}

	// 타자 기록 생성
	_, err = h.client.BatterRecord.Create().
		SetGames(1).
		SetAvg(0.400).
		SetPlateAppearances(5).
		SetAtBats(4).
		SetHits(2).
		SetRuns(1).
		SetHomeruns(1).
		SetRbis(2).
		SetUser(user).
		SetGame(game).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create batter record",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Test game and records created successfully",
		"game": game,
	})
}
