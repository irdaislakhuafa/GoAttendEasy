// Code generated by ent, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/irdaislakhuafa/GoAttendEasy/src/schema"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/attendance"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/reminder"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/role"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attendanceFields := schema.Attendance{}.Fields()
	_ = attendanceFields
	// attendanceDescUserID is the schema descriptor for user_id field.
	attendanceDescUserID := attendanceFields[0].Descriptor()
	// attendance.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	attendance.UserIDValidator = attendanceDescUserID.Validators[0].(func(string) error)
	// attendanceDescIsPresent is the schema descriptor for is_present field.
	attendanceDescIsPresent := attendanceFields[3].Descriptor()
	// attendance.DefaultIsPresent holds the default value on creation for the is_present field.
	attendance.DefaultIsPresent = attendanceDescIsPresent.Default.(bool)
	// attendanceDescCreatedAt is the schema descriptor for created_at field.
	attendanceDescCreatedAt := attendanceFields[5].Descriptor()
	// attendance.DefaultCreatedAt holds the default value on creation for the created_at field.
	attendance.DefaultCreatedAt = attendanceDescCreatedAt.Default.(time.Time)
	// attendanceDescCreatedBy is the schema descriptor for created_by field.
	attendanceDescCreatedBy := attendanceFields[6].Descriptor()
	// attendance.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	attendance.CreatedByValidator = attendanceDescCreatedBy.Validators[0].(func(string) error)
	// attendanceDescIsDeleted is the schema descriptor for is_deleted field.
	attendanceDescIsDeleted := attendanceFields[11].Descriptor()
	// attendance.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	attendance.DefaultIsDeleted = attendanceDescIsDeleted.Default.(bool)
	// attendanceDescID is the schema descriptor for id field.
	attendanceDescID := attendanceFields[4].Descriptor()
	// attendance.DefaultID holds the default value on creation for the id field.
	attendance.DefaultID = attendanceDescID.Default.(string)
	reminderFields := schema.Reminder{}.Fields()
	_ = reminderFields
	// reminderDescDay is the schema descriptor for day field.
	reminderDescDay := reminderFields[2].Descriptor()
	// reminder.DayValidator is a validator for the "day" field. It is called by the builders before save.
	reminder.DayValidator = reminderDescDay.Validators[0].(func(int) error)
	// reminderDescCreatedAt is the schema descriptor for created_at field.
	reminderDescCreatedAt := reminderFields[4].Descriptor()
	// reminder.DefaultCreatedAt holds the default value on creation for the created_at field.
	reminder.DefaultCreatedAt = reminderDescCreatedAt.Default.(time.Time)
	// reminderDescCreatedBy is the schema descriptor for created_by field.
	reminderDescCreatedBy := reminderFields[5].Descriptor()
	// reminder.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	reminder.CreatedByValidator = reminderDescCreatedBy.Validators[0].(func(string) error)
	// reminderDescIsDeleted is the schema descriptor for is_deleted field.
	reminderDescIsDeleted := reminderFields[10].Descriptor()
	// reminder.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	reminder.DefaultIsDeleted = reminderDescIsDeleted.Default.(bool)
	// reminderDescID is the schema descriptor for id field.
	reminderDescID := reminderFields[3].Descriptor()
	// reminder.DefaultID holds the default value on creation for the id field.
	reminder.DefaultID = reminderDescID.Default.(string)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescDescription is the schema descriptor for description field.
	roleDescDescription := roleFields[1].Descriptor()
	// role.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	role.DescriptionValidator = roleDescDescription.Validators[0].(func(string) error)
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[3].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(time.Time)
	// roleDescCreatedBy is the schema descriptor for created_by field.
	roleDescCreatedBy := roleFields[4].Descriptor()
	// role.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	role.CreatedByValidator = roleDescCreatedBy.Validators[0].(func(string) error)
	// roleDescIsDeleted is the schema descriptor for is_deleted field.
	roleDescIsDeleted := roleFields[9].Descriptor()
	// role.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	role.DefaultIsDeleted = roleDescIsDeleted.Default.(bool)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[2].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescRoleID is the schema descriptor for role_id field.
	userDescRoleID := userFields[3].Descriptor()
	// user.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	user.RoleIDValidator = userDescRoleID.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userFields[6].Descriptor()
	// user.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	user.CreatedByValidator = userDescCreatedBy.Validators[0].(func(string) error)
	// userDescIsDeleted is the schema descriptor for is_deleted field.
	userDescIsDeleted := userFields[11].Descriptor()
	// user.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	user.DefaultIsDeleted = userDescIsDeleted.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[4].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(string)
}
