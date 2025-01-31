package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// PitcherRecord holds the schema definition for the PitcherRecord entity.
type PitcherRecord struct {
    ent.Schema
}

// Fields of the PitcherRecord.
func (PitcherRecord) Fields() []ent.Field {
    return []ent.Field{
        field.Int("games").Default(0),
        field.Float("era").Optional(),
        field.Int("wins").Default(0),
        field.Int("losses").Default(0),
        field.Int("saves").Default(0),
        field.Int("holds").Default(0),
        field.Float("winning_pct").Optional(),
        field.Int("batters_faced").Default(0),
        field.Int("opponent_at_bats").Default(0),
        field.Float("innings").Default(0),
        field.Int("hits_allowed").Default(0),
        field.Int("homeruns_allowed").Default(0),
        field.Int("walks").Default(0),
        field.Int("hit_by_pitch").Default(0),
        field.Int("strikeouts").Default(0),
        field.Int("earned_runs").Default(0),
        field.Float("whip").Optional(),
        field.Float("opponent_avg").Optional(),
        field.Float("strikeout_rate").Optional(),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the PitcherRecord.
func (PitcherRecord) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).
            Ref("pitching_records").
            Unique().
            Required(),
        edge.From("game", Game.Type).
            Ref("pitching_records").
            Unique().
            Required(),
    }
}
