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

// CreateGameRequest는 새 게임 생성 요청을 위한 구조체입니다.
type CreateGameRequest struct {
    Date           string  `json:"date" validate:"required"`
    Opponent       *string `json:"opponent"`
    Location       *string `json:"location"`
    Result         *string `json:"result"`
    MyScore        *int    `json:"my_score"`
    OpponentScore  *int    `json:"opponent_score"`
}

// CreateBatterRecordRequest는 타자 기록 생성 요청을 위한 구조체입니다.
type CreateBatterRecordRequest struct {
    Game              CreateGameRequest `json:"game" validate:"required"`
    PlateAppearances  int              `json:"plate_appearances"`
    AtBats            int              `json:"at_bats"`
    Runs              int              `json:"runs"`
    Hits              int              `json:"hits"`
    Singles           int              `json:"singles"`
    Doubles           int              `json:"doubles"`
    Triples           int              `json:"triples"`
    Homeruns         int              `json:"homeruns"`
    Walks             int              `json:"walks"`
    Rbis              int              `json:"rbis"`
    Steals            int              `json:"steals"`
    HitByPitch        int              `json:"hit_by_pitch"`
    Strikeouts        int              `json:"strikeouts"`
    DoublePlays       int              `json:"double_plays"`
}

// CreatePitcherRecordRequest는 투수 기록 생성 요청을 위한 구조체입니다.
type CreatePitcherRecordRequest struct {
    Game             CreateGameRequest `json:"game" validate:"required"`
    Innings          float64          `json:"innings"`
    Wins             int              `json:"wins"`
    Losses           int              `json:"losses"`
    Saves            int              `json:"saves"`
    Holds            int              `json:"holds"`
    BattersFaced     int              `json:"batters_faced"`
    OpponentAtBats   int              `json:"opponent_at_bats"`
    HitsAllowed      int              `json:"hits_allowed"`
    HomerunsAllowed  int              `json:"homeruns_allowed"`
    Walks            int              `json:"walks"`
    HitByPitch       int              `json:"hit_by_pitch"`
    Strikeouts       int              `json:"strikeouts"`
    EarnedRuns       int              `json:"earned_runs"`
}

// CreateBatterRecord는 새로운 타자 기록을 생성합니다.
func (h *Handler) CreateBatterRecord(c *fiber.Ctx) error {
    // 사용자 ID 가져오기
    userID := c.Locals("userID").(int)

    // 요청 파싱
    var req CreateBatterRecordRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 유효성 검사
    if err := h.validator.Struct(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Validation failed",
        })
    }

    // 트랜잭션 시작
    tx, err := h.client.Tx(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to start transaction",
        })
    }

    // 게임 생성
    gameDate, err := time.Parse("2006-01-02", req.Game.Date)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid date format",
        })
    }

    game, err := tx.Game.Create().
        SetDate(gameDate).
        SetNillableOpponent(req.Game.Opponent).
        SetNillableLocation(req.Game.Location).
        SetNillableResult(req.Game.Result).
        SetNillableMyScore(req.Game.MyScore).
        SetNillableOpponentScore(req.Game.OpponentScore).
        SetUserID(userID).
        Save(c.Context())
    
    if err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create game",
        })
    }

    // 타율 계산
    avg := float64(0)
    if req.AtBats > 0 {
        avg = float64(req.Hits) / float64(req.AtBats)
    }

    // 장타율 계산
    slg := float64(0)
    if req.AtBats > 0 {
        totalBases := req.Singles + (req.Doubles * 2) + (req.Triples * 3) + (req.Homeruns * 4)
        slg = float64(totalBases) / float64(req.AtBats)
    }

    // 출루율 계산
    obp := float64(0)
    if (req.AtBats + req.Walks + req.HitByPitch) > 0 {
        obp = float64(req.Hits+req.Walks+req.HitByPitch) / float64(req.AtBats+req.Walks+req.HitByPitch)
    }

    // OPS 계산
    ops := slg + obp

    // BB/K 계산
    bbK := float64(0)
    if req.Strikeouts > 0 {
        bbK = float64(req.Walks) / float64(req.Strikeouts)
    }

    // 타자 기록 생성
    record, err := tx.BatterRecord.Create().
        SetUserID(userID).
        SetGame(game).
        SetGames(1).
        SetPlateAppearances(req.PlateAppearances).
        SetAtBats(req.AtBats).
        SetRuns(req.Runs).
        SetHits(req.Hits).
        SetSingles(req.Singles).
        SetDoubles(req.Doubles).
        SetTriples(req.Triples).
        SetHomeruns(req.Homeruns).
        SetWalks(req.Walks).
        SetRbis(req.Rbis).
        SetSteals(req.Steals).
        SetHitByPitch(req.HitByPitch).
        SetStrikeouts(req.Strikeouts).
        SetDoublePlays(req.DoublePlays).
        SetAvg(avg).
        SetSlg(slg).
        SetObp(obp).
        SetOps(ops).
        SetBbK(bbK).
        Save(c.Context())

    if err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create batter record",
        })
    }

    // 트랜잭션 커밋
    if err := tx.Commit(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to commit transaction",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Record created successfully",
        "record": record,
    })
}

// CreatePitcherRecord는 새로운 투수 기록을 생성합니다.
func (h *Handler) CreatePitcherRecord(c *fiber.Ctx) error {
    // 사용자 ID 가져오기
    userID := c.Locals("userID").(int)

    // 요청 파싱
    var req CreatePitcherRecordRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 유효성 검사
    if err := h.validator.Struct(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Validation failed",
        })
    }

    // 트랜잭션 시작
    tx, err := h.client.Tx(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to start transaction",
        })
    }

    // 게임 생성
    gameDate, err := time.Parse("2006-01-02", req.Game.Date)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid date format",
        })
    }

    game, err := tx.Game.Create().
        SetDate(gameDate).
        SetNillableOpponent(req.Game.Opponent).
        SetNillableLocation(req.Game.Location).
        SetNillableResult(req.Game.Result).
        SetNillableMyScore(req.Game.MyScore).
        SetNillableOpponentScore(req.Game.OpponentScore).
        SetUserID(userID).
        Save(c.Context())
    
    if err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create game",
        })
    }

    // ERA 계산
    era := float64(0)
    if req.Innings > 0 {
        era = (float64(req.EarnedRuns) * 9) / req.Innings
    }

    // WHIP 계산
    whip := float64(0)
    if req.Innings > 0 {
        whip = float64(req.Walks+req.HitsAllowed) / req.Innings
    }

    // 승률 계산
    winningPct := float64(0)
    if (req.Wins + req.Losses) > 0 {
        winningPct = float64(req.Wins) / float64(req.Wins+req.Losses)
    }

    // 삼진율 계산
    strikeoutRate := float64(0)
    if req.BattersFaced > 0 {
        strikeoutRate = float64(req.Strikeouts) / float64(req.BattersFaced)
    }

    // 피안타율 계산
    opponentAvg := float64(0)
    if req.OpponentAtBats > 0 {
        opponentAvg = float64(req.HitsAllowed) / float64(req.OpponentAtBats)
    }

    // 투수 기록 생성
    record, err := tx.PitcherRecord.Create().
        SetUserID(userID).
        SetGame(game).
        SetGames(1).
        SetInnings(req.Innings).
        SetWins(req.Wins).
        SetLosses(req.Losses).
        SetSaves(req.Saves).
        SetHolds(req.Holds).
        SetBattersFaced(req.BattersFaced).
        SetOpponentAtBats(req.OpponentAtBats).
        SetHitsAllowed(req.HitsAllowed).
        SetHomerunsAllowed(req.HomerunsAllowed).
        SetWalks(req.Walks).
        SetHitByPitch(req.HitByPitch).
        SetStrikeouts(req.Strikeouts).
        SetEarnedRuns(req.EarnedRuns).
        SetEra(era).
        SetWhip(whip).
        SetWinningPct(winningPct).
        SetStrikeoutRate(strikeoutRate).
        SetOpponentAvg(opponentAvg).
        Save(c.Context())

    if err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create pitcher record",
        })
    }

    // 트랜잭션 커밋
    if err := tx.Commit(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to commit transaction",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Record created successfully",
        "record": record,
    })
}
