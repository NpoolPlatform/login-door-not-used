// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/login-door/pkg/db/ent/loginrecord"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/provider"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	loginrecordFields := schema.LoginRecord{}.Fields()
	_ = loginrecordFields
	// loginrecordDescLoginTime is the schema descriptor for login_time field.
	loginrecordDescLoginTime := loginrecordFields[3].Descriptor()
	// loginrecord.DefaultLoginTime holds the default value on creation for the login_time field.
	loginrecord.DefaultLoginTime = loginrecordDescLoginTime.Default.(func() uint32)
	// loginrecordDescID is the schema descriptor for id field.
	loginrecordDescID := loginrecordFields[0].Descriptor()
	// loginrecord.DefaultID holds the default value on creation for the id field.
	loginrecord.DefaultID = loginrecordDescID.Default.(func() uuid.UUID)
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescCreateAt is the schema descriptor for create_at field.
	providerDescCreateAt := providerFields[6].Descriptor()
	// provider.DefaultCreateAt holds the default value on creation for the create_at field.
	provider.DefaultCreateAt = providerDescCreateAt.Default.(func() uint32)
	// providerDescUpdateAt is the schema descriptor for update_at field.
	providerDescUpdateAt := providerFields[7].Descriptor()
	// provider.DefaultUpdateAt holds the default value on creation for the update_at field.
	provider.DefaultUpdateAt = providerDescUpdateAt.Default.(func() uint32)
	// provider.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	provider.UpdateDefaultUpdateAt = providerDescUpdateAt.UpdateDefault.(func() uint32)
	// providerDescDeleteAt is the schema descriptor for delete_at field.
	providerDescDeleteAt := providerFields[8].Descriptor()
	// provider.DefaultDeleteAt holds the default value on creation for the delete_at field.
	provider.DefaultDeleteAt = providerDescDeleteAt.Default.(func() uint32)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() uuid.UUID)
}
