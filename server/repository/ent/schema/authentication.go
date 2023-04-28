package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Authentication holds the schema definition for the Authentication entity.
type Authentication struct {
	ent.Schema
}

// Fields of the Authentication.
func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().StorageKey("username").Comment("登陆名"),
		field.String("name").StorageKey("name").Comment("姓名"),
		field.Bytes("password").StorageKey("password").Comment("密码"),
	}
}

// Edges of the Authentication.
func (Authentication) Edges() []ent.Edge {
	return nil
}
