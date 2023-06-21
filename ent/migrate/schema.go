// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DefinitionsColumns holds the columns for the "definitions" table.
	DefinitionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString},
		{Name: "user_definitions", Type: field.TypeInt, Nullable: true},
		{Name: "word_definitions", Type: field.TypeInt, Nullable: true},
	}
	// DefinitionsTable holds the schema information for the "definitions" table.
	DefinitionsTable = &schema.Table{
		Name:       "definitions",
		Columns:    DefinitionsColumns,
		PrimaryKey: []*schema.Column{DefinitionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "definitions_users_definitions",
				Columns:    []*schema.Column{DefinitionsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "definitions_words_definitions",
				Columns:    []*schema.Column{DefinitionsColumns[3]},
				RefColumns: []*schema.Column{WordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "auth_id", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WordsColumns holds the columns for the "words" table.
	WordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString},
		{Name: "group_words", Type: field.TypeInt, Nullable: true},
		{Name: "user_words", Type: field.TypeInt, Nullable: true},
	}
	// WordsTable holds the schema information for the "words" table.
	WordsTable = &schema.Table{
		Name:       "words",
		Columns:    WordsColumns,
		PrimaryKey: []*schema.Column{WordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "words_groups_words",
				Columns:    []*schema.Column{WordsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "words_users_words",
				Columns:    []*schema.Column{WordsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DefinitionsTable,
		GroupsTable,
		UsersTable,
		WordsTable,
		UserGroupsTable,
	}
)

func init() {
	DefinitionsTable.ForeignKeys[0].RefTable = UsersTable
	DefinitionsTable.ForeignKeys[1].RefTable = WordsTable
	WordsTable.ForeignKeys[0].RefTable = GroupsTable
	WordsTable.ForeignKeys[1].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
