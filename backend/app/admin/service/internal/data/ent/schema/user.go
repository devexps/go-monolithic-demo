package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the user
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		//entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("id").
			Unique().
			MaxLen(36),
		field.String("user_name").
			Comment("user_name").
			Unique().
			MaxLen(50).
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),
		field.String("password").
			Comment("password").
			MaxLen(255),
		field.String("nick_name").
			Comment("nick_name").
			MaxLen(128).
			Optional(),
		field.String("real_name").
			Comment("real_name").
			MaxLen(128).
			Optional(),
		field.String("email").
			Comment("email").
			MaxLen(127).
			Optional(),
		field.String("phone").
			Comment("phone").
			MaxLen(11).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("id", "user_name").Unique(),
	}
}
