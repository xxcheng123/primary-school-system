package schema

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:35:33
 * @LastEditTime: 2024-03-08 15:18:02
 */

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xxcheng123/primary-school-system/ent/schema/mixins"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("parent_id").
			Default(0),
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.Int64("status").
			Default(1),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type),
	}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SnowflakeIDMixin{},
		mixin.Time{},
		mixins.SoftDeleteMixin{},
	}
}
