package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/bwmarrin/snowflake"
)

var node, err = snowflake.NewNode(1)

func genID() int64 {
	return node.Generate().Int64()
}

type SnowflakeIDMixin struct {
	mixin.Schema
	mixin.Time
}

func (SnowflakeIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").DefaultFunc(genID).Comment("雪花ID"),
	}
}
