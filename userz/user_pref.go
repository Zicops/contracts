package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_preferences (
// 	id varchar,
// 	user_id varchar,
// 	user_lsp_id varchar,
// 	sub_cateogory varchar,
// 	is_base boolean,
// 	is_active boolean,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserPreferencesTableMeta = table.Metadata{
	Name: "user_preferences",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"sub_category",
		"is_base",
		"is_active",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"user_id",
		"id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserPreferencesTable = table.New(UserPreferencesTableMeta)

type UserPreferences struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	UserLspID   string `db:"user_lsp_id"`
	SubCategory string `db:"sub_category"`
	IsBase      bool   `db:"is_base"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
}
