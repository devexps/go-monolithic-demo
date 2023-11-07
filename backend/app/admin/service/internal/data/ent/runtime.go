// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/schema"
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for userName field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "userName" field. It is called by the builders before save.
	user.UserNameValidator = func() func(string) error {
		validators := userDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(userName string) error {
			for _, fn := range fns {
				if err := fn(userName); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescNickName is the schema descriptor for nickName field.
	userDescNickName := userFields[3].Descriptor()
	// user.NickNameValidator is a validator for the "nickName" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescRealName is the schema descriptor for realName field.
	userDescRealName := userFields[4].Descriptor()
	// user.RealNameValidator is a validator for the "realName" field. It is called by the builders before save.
	user.RealNameValidator = userDescRealName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[6].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}