// Code generated by ent, DO NOT EDIT.

package attendance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUserID, v))
}

// In applies equality check predicate on the "in" field. It's identical to InEQ.
func In(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIn, v))
}

// Out applies equality check predicate on the "out" field. It's identical to OutEQ.
func Out(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldOut, v))
}

// IsPresent applies equality check predicate on the "is_present" field. It's identical to IsPresentEQ.
func IsPresent(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsPresent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedBy, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsDeleted, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContainsFold(FieldUserID, v))
}

// InEQ applies the EQ predicate on the "in" field.
func InEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIn, v))
}

// InNEQ applies the NEQ predicate on the "in" field.
func InNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldIn, v))
}

// InIn applies the In predicate on the "in" field.
func InIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldIn, vs...))
}

// InNotIn applies the NotIn predicate on the "in" field.
func InNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldIn, vs...))
}

// InGT applies the GT predicate on the "in" field.
func InGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldIn, v))
}

// InGTE applies the GTE predicate on the "in" field.
func InGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldIn, v))
}

// InLT applies the LT predicate on the "in" field.
func InLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldIn, v))
}

// InLTE applies the LTE predicate on the "in" field.
func InLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldIn, v))
}

// OutEQ applies the EQ predicate on the "out" field.
func OutEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldOut, v))
}

// OutNEQ applies the NEQ predicate on the "out" field.
func OutNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldOut, v))
}

// OutIn applies the In predicate on the "out" field.
func OutIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldOut, vs...))
}

// OutNotIn applies the NotIn predicate on the "out" field.
func OutNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldOut, vs...))
}

// OutGT applies the GT predicate on the "out" field.
func OutGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldOut, v))
}

// OutGTE applies the GTE predicate on the "out" field.
func OutGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldOut, v))
}

// OutLT applies the LT predicate on the "out" field.
func OutLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldOut, v))
}

// OutLTE applies the LTE predicate on the "out" field.
func OutLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldOut, v))
}

// IsPresentEQ applies the EQ predicate on the "is_present" field.
func IsPresentEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsPresent, v))
}

// IsPresentNEQ applies the NEQ predicate on the "is_present" field.
func IsPresentNEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldIsPresent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContainsFold(FieldDeletedBy, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldIsDeleted, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(sql.NotPredicates(p))
}