package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "regexp"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("email").
            Unique().
            NotEmpty().
            Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")),
        field.String("password").
            NotEmpty().
            Sensitive(),
        field.String("name").
            NotEmpty(),
        field.String("team").
            Optional(),
        field.Time("created_at").
            Default(time.Now).
            Immutable(),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("games", Game.Type),
        edge.To("batting_records", BatterRecord.Type),
        edge.To("pitching_records", PitcherRecord.Type),
    }
}
