package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// BatterRecord holds the schema definition for the BatterRecord entity.
type BatterRecord struct {
    ent.Schema
}

// Fields of the BatterRecord.
func (BatterRecord) Fields() []ent.Field {
    return []ent.Field{
        field.Int("games").Default(0),
        field.Float("avg").Optional(),
        field.Int("plate_appearances").Default(0),
        field.Int("at_bats").Default(0),
        field.Int("runs").Default(0),
        field.Int("hits").Default(0),
        field.Int("singles").Default(0),
        field.Int("doubles").Default(0),
        field.Int("triples").Default(0),
        field.Int("homeruns").Default(0),
        field.Int("walks").Default(0),
        field.Int("rbis").Default(0),
        field.Int("steals").Default(0),
        field.Int("hit_by_pitch").Default(0),
        field.Int("strikeouts").Default(0),
        field.Int("double_plays").Default(0),
        field.Float("slg").Optional(),
        field.Float("obp").Optional(),
        field.Float("ops").Optional(),
        field.Float("bb_k").Optional(),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the BatterRecord.
func (BatterRecord) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).
            Ref("batting_records").
            Unique().
            Required(),
        edge.From("game", Game.Type).
            Ref("batting_records").
            Unique().
            Required(),
    }
}
