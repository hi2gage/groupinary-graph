// Code generated by ent, DO NOT EDIT.

package ent

import (
	"groupinary/ent/definition"
	"groupinary/ent/group"
	"groupinary/ent/schema"
	"groupinary/ent/user"
	"groupinary/ent/word"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	definitionMixin := schema.Definition{}.Mixin()
	definitionMixinFields0 := definitionMixin[0].Fields()
	_ = definitionMixinFields0
	definitionFields := schema.Definition{}.Fields()
	_ = definitionFields
	// definitionDescCreateTime is the schema descriptor for create_time field.
	definitionDescCreateTime := definitionMixinFields0[0].Descriptor()
	// definition.DefaultCreateTime holds the default value on creation for the create_time field.
	definition.DefaultCreateTime = definitionDescCreateTime.Default.(func() time.Time)
	// definitionDescUpdateTime is the schema descriptor for update_time field.
	definitionDescUpdateTime := definitionMixinFields0[1].Descriptor()
	// definition.DefaultUpdateTime holds the default value on creation for the update_time field.
	definition.DefaultUpdateTime = definitionDescUpdateTime.Default.(func() time.Time)
	// definition.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	definition.UpdateDefaultUpdateTime = definitionDescUpdateTime.UpdateDefault.(func() time.Time)
	// definitionDescDescription is the schema descriptor for description field.
	definitionDescDescription := definitionFields[0].Descriptor()
	// definition.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	definition.DescriptionValidator = definitionDescDescription.Validators[0].(func(string) error)
	groupMixin := schema.Group{}.Mixin()
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreateTime is the schema descriptor for create_time field.
	groupDescCreateTime := groupMixinFields0[0].Descriptor()
	// group.DefaultCreateTime holds the default value on creation for the create_time field.
	group.DefaultCreateTime = groupDescCreateTime.Default.(func() time.Time)
	// groupDescUpdateTime is the schema descriptor for update_time field.
	groupDescUpdateTime := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdateTime holds the default value on creation for the update_time field.
	group.DefaultUpdateTime = groupDescUpdateTime.Default.(func() time.Time)
	// group.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	group.UpdateDefaultUpdateTime = groupDescUpdateTime.UpdateDefault.(func() time.Time)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescAuthID is the schema descriptor for authID field.
	userDescAuthID := userFields[0].Descriptor()
	// user.AuthIDValidator is a validator for the "authID" field. It is called by the builders before save.
	user.AuthIDValidator = userDescAuthID.Validators[0].(func(string) error)
	wordMixin := schema.Word{}.Mixin()
	wordMixinFields0 := wordMixin[0].Fields()
	_ = wordMixinFields0
	wordFields := schema.Word{}.Fields()
	_ = wordFields
	// wordDescCreateTime is the schema descriptor for create_time field.
	wordDescCreateTime := wordMixinFields0[0].Descriptor()
	// word.DefaultCreateTime holds the default value on creation for the create_time field.
	word.DefaultCreateTime = wordDescCreateTime.Default.(func() time.Time)
	// wordDescUpdateTime is the schema descriptor for update_time field.
	wordDescUpdateTime := wordMixinFields0[1].Descriptor()
	// word.DefaultUpdateTime holds the default value on creation for the update_time field.
	word.DefaultUpdateTime = wordDescUpdateTime.Default.(func() time.Time)
	// word.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	word.UpdateDefaultUpdateTime = wordDescUpdateTime.UpdateDefault.(func() time.Time)
	// wordDescDescription is the schema descriptor for description field.
	wordDescDescription := wordFields[0].Descriptor()
	// word.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	word.DescriptionValidator = wordDescDescription.Validators[0].(func(string) error)
}
