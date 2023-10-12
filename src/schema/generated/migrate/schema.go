// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttendancesColumns holds the columns for the "attendances" table.
	AttendancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "in", Type: field.TypeTime},
		{Name: "out", Type: field.TypeTime},
		{Name: "is_present", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// AttendancesTable holds the schema information for the "attendances" table.
	AttendancesTable = &schema.Table{
		Name:       "attendances",
		Columns:    AttendancesColumns,
		PrimaryKey: []*schema.Column{AttendancesColumns[0]},
	}
	// RemindersColumns holds the columns for the "reminders" table.
	RemindersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "in", Type: field.TypeTime},
		{Name: "out", Type: field.TypeTime},
		{Name: "day", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// RemindersTable holds the schema information for the "reminders" table.
	RemindersTable = &schema.Table{
		Name:       "reminders",
		Columns:    RemindersColumns,
		PrimaryKey: []*schema.Column{RemindersColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttendancesTable,
		RemindersTable,
		RolesTable,
		UsersTable,
	}
)

func init() {
}
