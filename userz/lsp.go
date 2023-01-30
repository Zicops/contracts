package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.learning_space(
//	id varchar,
//	org_id varchar,
//	org_unit_id varchar,
//  type varchar,
//	name varchar,
//	logo_bucket varchar,
//	logo_url varchar,
//	profile_picture_bucket varchar,
//	profile_picture_url varchar,
//	no_of_users bigint,
//	owners set<varchar>,
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
		"type",
		"name",
		"logo_bucket",
		"logo_url",
		"profile_picture_bucket",
		"profile_picture_url",
		"no_of_users",
		"owners",
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
	ID                   string   `db:"id"`
	OrgID                string   `db:"org_id"`
	OrgUnitID            string   `db:"org_unit_id"`
	Type                 string   `db:"type"`
	Name                 string   `db:"name"`
	LogoBucket           string   `db:"logo_bucket"`
	LogoURL              string   `db:"logo_url"`
	ProfilePictureBucket string   `db:"profile_picture_bucket"`
	ProfilePictureURL    string   `db:"profile_picture_url"`
	NoOfUsers            int64    `db:"no_of_users"`
	Owners               []string `db:"owners"`
	IsDefault            bool     `db:"is_default"`
	Status               string   `db:"status"`
	CreatedAt            int64    `db:"created_at"`
	CreatedBy            string   `db:"created_by"`
	UpdatedAt            int64    `db:"updated_at"`
	UpdatedBy            string   `db:"updated_by"`
}
