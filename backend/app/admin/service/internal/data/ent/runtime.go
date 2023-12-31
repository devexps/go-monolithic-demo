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
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields3 := userMixin[3].Fields()
	_ = userMixinFields3
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[2].Descriptor()
	// user.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescRealName is the schema descriptor for real_name field.
	userDescRealName := userFields[3].Descriptor()
	// user.RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	user.RealNameValidator = userDescRealName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[5].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint32) error)
}
