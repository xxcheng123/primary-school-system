// Code generated by ent, DO NOT EDIT.

package swiper

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/xxcheng123/primary-school-system/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldUpdateTime, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldDeletedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldTitle, v))
}

// Img applies equality check predicate on the "img" field. It's identical to ImgEQ.
func Img(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldImg, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldURL, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldOrder, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldUpdateTime, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Swiper {
	return predicate.Swiper(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Swiper {
	return predicate.Swiper(sql.FieldNotNull(FieldDeletedAt))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContainsFold(FieldTitle, v))
}

// ImgEQ applies the EQ predicate on the "img" field.
func ImgEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldImg, v))
}

// ImgNEQ applies the NEQ predicate on the "img" field.
func ImgNEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldImg, v))
}

// ImgIn applies the In predicate on the "img" field.
func ImgIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldImg, vs...))
}

// ImgNotIn applies the NotIn predicate on the "img" field.
func ImgNotIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldImg, vs...))
}

// ImgGT applies the GT predicate on the "img" field.
func ImgGT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldImg, v))
}

// ImgGTE applies the GTE predicate on the "img" field.
func ImgGTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldImg, v))
}

// ImgLT applies the LT predicate on the "img" field.
func ImgLT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldImg, v))
}

// ImgLTE applies the LTE predicate on the "img" field.
func ImgLTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldImg, v))
}

// ImgContains applies the Contains predicate on the "img" field.
func ImgContains(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContains(FieldImg, v))
}

// ImgHasPrefix applies the HasPrefix predicate on the "img" field.
func ImgHasPrefix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasPrefix(FieldImg, v))
}

// ImgHasSuffix applies the HasSuffix predicate on the "img" field.
func ImgHasSuffix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasSuffix(FieldImg, v))
}

// ImgEqualFold applies the EqualFold predicate on the "img" field.
func ImgEqualFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEqualFold(FieldImg, v))
}

// ImgContainsFold applies the ContainsFold predicate on the "img" field.
func ImgContainsFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContainsFold(FieldImg, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Swiper {
	return predicate.Swiper(sql.FieldContainsFold(FieldURL, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldOrder, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Swiper {
	return predicate.Swiper(sql.FieldLTE(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Swiper) predicate.Swiper {
	return predicate.Swiper(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Swiper) predicate.Swiper {
	return predicate.Swiper(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Swiper) predicate.Swiper {
	return predicate.Swiper(sql.NotPredicates(p))
}