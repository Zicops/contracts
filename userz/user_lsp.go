package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_lsp_map (
//	id varchar,
//	user_id varchar,
//	lsp_id varchar,
//	status varchar,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );

var UserLSPMapTableMeta = table.Metadata{
	Name: "user_lsp_map",
	Columns: []string{
		"id",
		"user_id",
		"lsp_id",
		"status",
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

var UserLspTable = table.New(UserLSPMapTableMeta)

type UserLsp struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	LspID     string `db:"lsp_id"`
	Status    string `db:"status"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
}