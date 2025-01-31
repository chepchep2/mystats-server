package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Game holds the schema definition for the Game entity.
type Game struct {
    ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
    return []ent.Field{
        field.Time("date"),
        field.String("opponent").Optional(),
        field.String("location").Optional(),
        field.String("result").Optional(),
        field.Int("my_score").Optional(),
        field.Int("opponent_score").Optional(),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).
            Ref("games").
            Unique().
            Required(),
        edge.To("batting_records", BatterRecord.Type),
        edge.To("pitching_records", PitcherRecord.Type),
    }
}
