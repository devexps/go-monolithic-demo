// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldRealName holds the string denoting the real_name field in the database.
	FieldRealName = "real_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldAuthority holds the string denoting the authority field in the database.
	FieldAuthority = "authority"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateBy,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldStatus,
	FieldUsername,
	FieldPassword,
	FieldNickName,
	FieldRealName,
	FieldEmail,
	FieldPhone,
	FieldAuthority,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	NickNameValidator func(string) error
	// RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	RealNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusON is the default value of the Status enum.
const DefaultStatus = StatusON

// Status values.
const (
	StatusOFF Status = "OFF"
	StatusON  Status = "ON"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOFF, StatusON:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Authority defines the type for the "authority" enum field.
type Authority string

// AuthorityCUSTOMER_USER is the default value of the Authority enum.
const DefaultAuthority = AuthorityCUSTOMER_USER

// Authority values.
const (
	AuthorityUA_UNSPECIFIED Authority = "UA_UNSPECIFIED"
	AuthoritySYS_ADMIN      Authority = "SYS_ADMIN"
	AuthorityCUSTOMER_USER  Authority = "CUSTOMER_USER"
	AuthorityGUEST_USER     Authority = "GUEST_USER"
	AuthorityREFRESH_TOKEN  Authority = "REFRESH_TOKEN"
)

func (a Authority) String() string {
	return string(a)
}

// AuthorityValidator is a validator for the "authority" field enum values. It is called by the builders before save.
func AuthorityValidator(a Authority) error {
	switch a {
	case AuthorityUA_UNSPECIFIED, AuthoritySYS_ADMIN, AuthorityCUSTOMER_USER, AuthorityGUEST_USER, AuthorityREFRESH_TOKEN:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for authority field: %q", a)
	}
}
