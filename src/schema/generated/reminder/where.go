// Code generated by ent, DO NOT EDIT.

package reminder

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContainsFold(FieldID, id))
}

// In applies equality check predicate on the "in" field. It's identical to InEQ.
func In(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldIn, v))
}

// Out applies equality check predicate on the "out" field. It's identical to OutEQ.
func Out(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldOut, v))
}

// Day applies equality check predicate on the "day" field. It's identical to DayEQ.
func Day(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDay, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDeletedBy, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldIsDeleted, v))
}

// InEQ applies the EQ predicate on the "in" field.
func InEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldIn, v))
}

// InNEQ applies the NEQ predicate on the "in" field.
func InNEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldIn, v))
}

// InIn applies the In predicate on the "in" field.
func InIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldIn, vs...))
}

// InNotIn applies the NotIn predicate on the "in" field.
func InNotIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldIn, vs...))
}

// InGT applies the GT predicate on the "in" field.
func InGT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldIn, v))
}

// InGTE applies the GTE predicate on the "in" field.
func InGTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldIn, v))
}

// InLT applies the LT predicate on the "in" field.
func InLT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldIn, v))
}

// InLTE applies the LTE predicate on the "in" field.
func InLTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldIn, v))
}

// OutEQ applies the EQ predicate on the "out" field.
func OutEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldOut, v))
}

// OutNEQ applies the NEQ predicate on the "out" field.
func OutNEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldOut, v))
}

// OutIn applies the In predicate on the "out" field.
func OutIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldOut, vs...))
}

// OutNotIn applies the NotIn predicate on the "out" field.
func OutNotIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldOut, vs...))
}

// OutGT applies the GT predicate on the "out" field.
func OutGT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldOut, v))
}

// OutGTE applies the GTE predicate on the "out" field.
func OutGTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldOut, v))
}

// OutLT applies the LT predicate on the "out" field.
func OutLT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldOut, v))
}

// OutLTE applies the LTE predicate on the "out" field.
func OutLTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldOut, v))
}

// DayEQ applies the EQ predicate on the "day" field.
func DayEQ(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDay, v))
}

// DayNEQ applies the NEQ predicate on the "day" field.
func DayNEQ(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldDay, v))
}

// DayIn applies the In predicate on the "day" field.
func DayIn(vs ...int) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldDay, vs...))
}

// DayNotIn applies the NotIn predicate on the "day" field.
func DayNotIn(vs ...int) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldDay, vs...))
}

// DayGT applies the GT predicate on the "day" field.
func DayGT(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldDay, v))
}

// DayGTE applies the GTE predicate on the "day" field.
func DayGTE(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldDay, v))
}

// DayLT applies the LT predicate on the "day" field.
func DayLT(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldDay, v))
}

// DayLTE applies the LTE predicate on the "day" field.
func DayLTE(v int) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldDay, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldNotNull(FieldCreatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Reminder {
	return predicate.Reminder(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Reminder {
	return predicate.Reminder(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Reminder {
	return predicate.Reminder(sql.FieldContainsFold(FieldDeletedBy, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Reminder {
	return predicate.Reminder(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Reminder {
	return predicate.Reminder(sql.FieldNEQ(FieldIsDeleted, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reminder) predicate.Reminder {
	return predicate.Reminder(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reminder) predicate.Reminder {
	return predicate.Reminder(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Reminder) predicate.Reminder {
	return predicate.Reminder(sql.NotPredicates(p))
}
