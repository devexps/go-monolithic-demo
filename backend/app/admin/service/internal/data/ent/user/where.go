// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// CreateBy applies equality check predicate on the "create_by" field. It's identical to CreateByEQ.
func CreateBy(v uint32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateBy, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// RealName applies equality check predicate on the "real_name" field. It's identical to RealNameEQ.
func RealName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// CreateByEQ applies the EQ predicate on the "create_by" field.
func CreateByEQ(v uint32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateBy, v))
}

// CreateByNEQ applies the NEQ predicate on the "create_by" field.
func CreateByNEQ(v uint32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateBy, v))
}

// CreateByIn applies the In predicate on the "create_by" field.
func CreateByIn(vs ...uint32) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateBy, vs...))
}

// CreateByNotIn applies the NotIn predicate on the "create_by" field.
func CreateByNotIn(vs ...uint32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateBy, vs...))
}

// CreateByGT applies the GT predicate on the "create_by" field.
func CreateByGT(v uint32) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateBy, v))
}

// CreateByGTE applies the GTE predicate on the "create_by" field.
func CreateByGTE(v uint32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateBy, v))
}

// CreateByLT applies the LT predicate on the "create_by" field.
func CreateByLT(v uint32) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateBy, v))
}

// CreateByLTE applies the LTE predicate on the "create_by" field.
func CreateByLTE(v uint32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateBy, v))
}

// CreateByIsNil applies the IsNil predicate on the "create_by" field.
func CreateByIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCreateBy))
}

// CreateByNotNil applies the NotNil predicate on the "create_by" field.
func CreateByNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCreateBy))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeleteTime))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldStatus))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameIsNil applies the IsNil predicate on the "nick_name" field.
func NickNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldNickName))
}

// NickNameNotNil applies the NotNil predicate on the "nick_name" field.
func NickNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldNickName))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickName, v))
}

// RealNameEQ applies the EQ predicate on the "real_name" field.
func RealNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// RealNameNEQ applies the NEQ predicate on the "real_name" field.
func RealNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRealName, v))
}

// RealNameIn applies the In predicate on the "real_name" field.
func RealNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRealName, vs...))
}

// RealNameNotIn applies the NotIn predicate on the "real_name" field.
func RealNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRealName, vs...))
}

// RealNameGT applies the GT predicate on the "real_name" field.
func RealNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRealName, v))
}

// RealNameGTE applies the GTE predicate on the "real_name" field.
func RealNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRealName, v))
}

// RealNameLT applies the LT predicate on the "real_name" field.
func RealNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRealName, v))
}

// RealNameLTE applies the LTE predicate on the "real_name" field.
func RealNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRealName, v))
}

// RealNameContains applies the Contains predicate on the "real_name" field.
func RealNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRealName, v))
}

// RealNameHasPrefix applies the HasPrefix predicate on the "real_name" field.
func RealNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRealName, v))
}

// RealNameHasSuffix applies the HasSuffix predicate on the "real_name" field.
func RealNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRealName, v))
}

// RealNameIsNil applies the IsNil predicate on the "real_name" field.
func RealNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldRealName))
}

// RealNameNotNil applies the NotNil predicate on the "real_name" field.
func RealNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldRealName))
}

// RealNameEqualFold applies the EqualFold predicate on the "real_name" field.
func RealNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRealName, v))
}

// RealNameContainsFold applies the ContainsFold predicate on the "real_name" field.
func RealNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRealName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
