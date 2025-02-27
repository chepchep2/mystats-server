// Code generated by ent, DO NOT EDIT.

package pitcherrecord

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pitcherrecord type in the database.
	Label = "pitcher_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGames holds the string denoting the games field in the database.
	FieldGames = "games"
	// FieldEra holds the string denoting the era field in the database.
	FieldEra = "era"
	// FieldWins holds the string denoting the wins field in the database.
	FieldWins = "wins"
	// FieldLosses holds the string denoting the losses field in the database.
	FieldLosses = "losses"
	// FieldSaves holds the string denoting the saves field in the database.
	FieldSaves = "saves"
	// FieldHolds holds the string denoting the holds field in the database.
	FieldHolds = "holds"
	// FieldWinningPct holds the string denoting the winning_pct field in the database.
	FieldWinningPct = "winning_pct"
	// FieldBattersFaced holds the string denoting the batters_faced field in the database.
	FieldBattersFaced = "batters_faced"
	// FieldOpponentAtBats holds the string denoting the opponent_at_bats field in the database.
	FieldOpponentAtBats = "opponent_at_bats"
	// FieldInnings holds the string denoting the innings field in the database.
	FieldInnings = "innings"
	// FieldHitsAllowed holds the string denoting the hits_allowed field in the database.
	FieldHitsAllowed = "hits_allowed"
	// FieldHomerunsAllowed holds the string denoting the homeruns_allowed field in the database.
	FieldHomerunsAllowed = "homeruns_allowed"
	// FieldWalks holds the string denoting the walks field in the database.
	FieldWalks = "walks"
	// FieldHitByPitch holds the string denoting the hit_by_pitch field in the database.
	FieldHitByPitch = "hit_by_pitch"
	// FieldStrikeouts holds the string denoting the strikeouts field in the database.
	FieldStrikeouts = "strikeouts"
	// FieldEarnedRuns holds the string denoting the earned_runs field in the database.
	FieldEarnedRuns = "earned_runs"
	// FieldWhip holds the string denoting the whip field in the database.
	FieldWhip = "whip"
	// FieldOpponentAvg holds the string denoting the opponent_avg field in the database.
	FieldOpponentAvg = "opponent_avg"
	// FieldStrikeoutRate holds the string denoting the strikeout_rate field in the database.
	FieldStrikeoutRate = "strikeout_rate"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// Table holds the table name of the pitcherrecord in the database.
	Table = "pitcher_records"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "pitcher_records"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_pitching_records"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "pitcher_records"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_pitching_records"
)

// Columns holds all SQL columns for pitcherrecord fields.
var Columns = []string{
	FieldID,
	FieldGames,
	FieldEra,
	FieldWins,
	FieldLosses,
	FieldSaves,
	FieldHolds,
	FieldWinningPct,
	FieldBattersFaced,
	FieldOpponentAtBats,
	FieldInnings,
	FieldHitsAllowed,
	FieldHomerunsAllowed,
	FieldWalks,
	FieldHitByPitch,
	FieldStrikeouts,
	FieldEarnedRuns,
	FieldWhip,
	FieldOpponentAvg,
	FieldStrikeoutRate,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pitcher_records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_pitching_records",
	"user_pitching_records",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultGames holds the default value on creation for the "games" field.
	DefaultGames int
	// DefaultWins holds the default value on creation for the "wins" field.
	DefaultWins int
	// DefaultLosses holds the default value on creation for the "losses" field.
	DefaultLosses int
	// DefaultSaves holds the default value on creation for the "saves" field.
	DefaultSaves int
	// DefaultHolds holds the default value on creation for the "holds" field.
	DefaultHolds int
	// DefaultBattersFaced holds the default value on creation for the "batters_faced" field.
	DefaultBattersFaced int
	// DefaultOpponentAtBats holds the default value on creation for the "opponent_at_bats" field.
	DefaultOpponentAtBats int
	// DefaultInnings holds the default value on creation for the "innings" field.
	DefaultInnings float64
	// DefaultHitsAllowed holds the default value on creation for the "hits_allowed" field.
	DefaultHitsAllowed int
	// DefaultHomerunsAllowed holds the default value on creation for the "homeruns_allowed" field.
	DefaultHomerunsAllowed int
	// DefaultWalks holds the default value on creation for the "walks" field.
	DefaultWalks int
	// DefaultHitByPitch holds the default value on creation for the "hit_by_pitch" field.
	DefaultHitByPitch int
	// DefaultStrikeouts holds the default value on creation for the "strikeouts" field.
	DefaultStrikeouts int
	// DefaultEarnedRuns holds the default value on creation for the "earned_runs" field.
	DefaultEarnedRuns int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the PitcherRecord queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGames orders the results by the games field.
func ByGames(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGames, opts...).ToFunc()
}

// ByEra orders the results by the era field.
func ByEra(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEra, opts...).ToFunc()
}

// ByWins orders the results by the wins field.
func ByWins(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWins, opts...).ToFunc()
}

// ByLosses orders the results by the losses field.
func ByLosses(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLosses, opts...).ToFunc()
}

// BySaves orders the results by the saves field.
func BySaves(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaves, opts...).ToFunc()
}

// ByHolds orders the results by the holds field.
func ByHolds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHolds, opts...).ToFunc()
}

// ByWinningPct orders the results by the winning_pct field.
func ByWinningPct(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWinningPct, opts...).ToFunc()
}

// ByBattersFaced orders the results by the batters_faced field.
func ByBattersFaced(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBattersFaced, opts...).ToFunc()
}

// ByOpponentAtBats orders the results by the opponent_at_bats field.
func ByOpponentAtBats(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpponentAtBats, opts...).ToFunc()
}

// ByInnings orders the results by the innings field.
func ByInnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInnings, opts...).ToFunc()
}

// ByHitsAllowed orders the results by the hits_allowed field.
func ByHitsAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHitsAllowed, opts...).ToFunc()
}

// ByHomerunsAllowed orders the results by the homeruns_allowed field.
func ByHomerunsAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHomerunsAllowed, opts...).ToFunc()
}

// ByWalks orders the results by the walks field.
func ByWalks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWalks, opts...).ToFunc()
}

// ByHitByPitch orders the results by the hit_by_pitch field.
func ByHitByPitch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHitByPitch, opts...).ToFunc()
}

// ByStrikeouts orders the results by the strikeouts field.
func ByStrikeouts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStrikeouts, opts...).ToFunc()
}

// ByEarnedRuns orders the results by the earned_runs field.
func ByEarnedRuns(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEarnedRuns, opts...).ToFunc()
}

// ByWhip orders the results by the whip field.
func ByWhip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhip, opts...).ToFunc()
}

// ByOpponentAvg orders the results by the opponent_avg field.
func ByOpponentAvg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpponentAvg, opts...).ToFunc()
}

// ByStrikeoutRate orders the results by the strikeout_rate field.
func ByStrikeoutRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStrikeoutRate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByGameField orders the results by game field.
func ByGameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newGameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GameTable, GameColumn),
	)
}
