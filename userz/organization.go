package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.organization (
//	id varchar,
//	name varchar,
//
//	logo_bucket varchar,
//	logo_url varchar,
//
//	industry varchar,
//	type varchar,
//	zicops_subdomain varchar,
//	emp_count text,
//	website varchar,
//	linkedin varchar,
//	facebook varchar,
//	twitter varchar,
//
//	status varchar,
//	created_at bigint,
//	created_by varchar,
//	updated_at bigint,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//
var OrganizationTableMeta = table.Metadata{
	Name: "organization",
	Columns: []string{
		"id",
		"name",
		"logo_bucket",
		"logo_url",
		"industry",
		"type",
		"zicops_subdomain",
		"emp_count",
		"website",
		"linkedin",
		"facebook",
		"twitter",
		"status",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

var OrganizationTable = table.New(OrganizationTableMeta)

type Organization struct {
	ID              string `db:"id"`
	Name            string `db:"name"`
	LogoBucket      string `db:"logo_bucket"`
	LogoURL         string `db:"logo_url"`
	Industry        string `db:"industry"`
	Type            string `db:"type"`
	ZicopsSubdomain string `db:"zicops_subdomain"`
	EmpCount        string `db:"emp_count"`
	Website         string `db:"website"`
	Linkedin        string `db:"linkedin"`
	Facebook        string `db:"facebook"`
	Twitter         string `db:"twitter"`
	Status          string `db:"status"`
	CreatedAt       int64  `db:"created_at"`
	CreatedBy       string `db:"created_by"`
	UpdatedAt       int64  `db:"updated_at"`
	UpdatedBy       string `db:"updated_by"`
}
