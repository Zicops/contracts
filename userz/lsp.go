package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.learning_space(
//	id varchar,
//	org_id varchar,
//	org_unit_id varchar,
//	lsp_name varchar,
//	lsp_logo_bucket varchar,
//	lsp_logo_url varchar,
//	lsp_profile_picture_bucket varchar,
//	lsp_profile_picture_url varchar,
//	no_of_users bigint,
//	lsp_owners set<varchar>,
//	is_default boolean,
//	status varchar,
//	created_at bigint,
//	created_by varchar,
//	updated_at bigint,
//	updated_by varchar,
//	PRIMARY KEY ((id, org_id), created_at)
//  )
//  WITH CLUSTERING ORDER BY (created_at DESC);

var LspTableMeta = table.Metadata{
	Name: "learning_space",
	Columns: []string{
		"id",
		"org_id",
		"org_unit_id",
		"lsp_name",
		"lsp_logo_bucket",
		"lsp_logo_url",
		"lsp_profile_picture_bucket",
		"lsp_profile_picture_url",
		"no_of_users",
		"lsp_owners",
		"is_default",
		"status",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	},
	PartKey: []string{
		"id",
		"org_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var LspTable = table.New(LspTableMeta)

type Lsp struct {
	ID                      string   `db:"id"`
	OrgID                   string   `db:"org_id"`
	OrgUnitID               string   `db:"org_unit_id"`
	LspName                 string   `db:"lsp_name"`
	LspLogoBucket           string   `db:"lsp_logo_bucket"`
	LspLogoURL              string   `db:"lsp_logo_url"`
	LspProfilePictureBucket string   `db:"lsp_profile_picture_bucket"`
	LspProfilePictureURL    string   `db:"lsp_profile_picture_url"`
	NoOfUsers               int64    `db:"no_of_users"`
	LspOwners               []string `db:"lsp_owners"`
	IsDefault               bool     `db:"is_default"`
	Status                  string   `db:"status"`
	CreatedAt               int64    `db:"created_at"`
	CreatedBy               string   `db:"created_by"`
	UpdatedAt               int64    `db:"updated_at"`
	UpdatedBy               string   `db:"updated_by"`
}
