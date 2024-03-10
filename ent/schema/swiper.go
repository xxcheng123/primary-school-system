package schema

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 17:35:12
 * @LastEditTime: 2024-03-08 17:38:48
 */

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xxcheng123/primary-school-system/ent/schema/mixins"
)

// Swiper holds the schema definition for the Swiper entity.
type Swiper struct {
	ent.Schema
}

// Fields of the Swiper.
func (Swiper) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default("默认标题"),
		field.String("img").Default("https://cdn-static.xxcheng.cn/static/blog/images/2023/04/24/cd46a916b71dcf8a4f38508c6b4f87f3.jpg"),
		field.String("url").Default("https://www.baidu.com/"),
		field.Int64("order").Default(1).Comment("越大优先级越高"),
		field.Int64("status").Default(1),
	}
}

// Edges of the Swiper.
func (Swiper) Edges() []ent.Edge {
	return nil
}
func (Swiper) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SnowflakeIDMixin{},
		mixin.Time{},
		mixins.SoftDeleteMixin{},
	}
}
