package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_role (
// 	id varchar,
// 	role varchar,
// 	user_id varchar,
// 	user_lsp_id varchar,
// 	is_active boolean,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserRoleMapTableMeta = table.Metadata{
	Name: "user_role",
	Columns: []string{
		"id",
		"role",
		"user_id",
		"user_lsp_id",
		"is_active",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

var UserRoleTable = table.New(UserRoleMapTableMeta)

type UserRole struct {
	ID        string `db:"id"`
	Role      string `db:"role"`
	UserID    string `db:"user_id"`
	UserLspID string `db:"user_lsp_id"`
	IsActive  bool   `db:"is_active"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
}
