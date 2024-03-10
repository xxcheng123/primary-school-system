package schema

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:34:25
 * @LastEditTime: 2024-03-08 14:34:26
 */

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xxcheng123/primary-school-system/ent/schema/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			MaxLen(255).
			NotEmpty(),
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("password").
			MaxLen(255).
			NotEmpty(),
		field.Int64("status").
			Default(1),
		field.Int64("role").Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SnowflakeIDMixin{},
		mixin.Time{},
		mixins.SoftDeleteMixin{},
	}
}
