package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xxcheng123/primary-school-system/ent/schema/mixins"
)

// XConfig holds the schema definition for the XConfig entity.
type XConfig struct {
	ent.Schema
}

// Fields of the XConfig.
func (XConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("value").Default("1"),
		field.String("desc").Default("default desc"),
	}
}

// Edges of the XConfig.
func (XConfig) Edges() []ent.Edge {
	return nil
}

func (XConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SnowflakeIDMixin{},
		mixin.Time{},
		mixins.SoftDeleteMixin{},
	}
}
