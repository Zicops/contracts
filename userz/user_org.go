package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_org_map (
//	id varchar,
//	user_id varchar,
//	user_lsp_id varchar,
//	org_id varchar,
//	org_role varchar,
//	is_active boolean,
//	emp_id varchar,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );

var UserOrgMapTableMeta = table.Metadata{
	Name: "user_org_map",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"org_id",
		"org_role",
		"is_active",
		"emp_id",
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

var UserOrgTable = table.New(UserOrgMapTableMeta)

type UserOrg struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	UserLspID string `db:"user_lsp_id"`
	OrgID     string `db:"org_id"`
	OrgRole   string `db:"org_role"`
	IsActive  bool   `db:"is_active"`
	EmpID     string `db:"emp_id"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
}
