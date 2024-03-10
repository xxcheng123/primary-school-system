package schema

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:35:47
 * @LastEditTime: 2024-03-08 15:18:06
 */

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xxcheng123/primary-school-system/ent/schema/mixins"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255).
			NotEmpty(),
		field.Text("content").
			NotEmpty(),
		field.Int64("status").
			Default(1),
		field.String("img").Default("https://zip.xxcheng.cn/lunwen/news.jpeg"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("article"),
		edge.From("category", Category.Type).Unique().Ref("article"),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SnowflakeIDMixin{},
		mixin.Time{},
		mixins.SoftDeleteMixin{},
	}
}
