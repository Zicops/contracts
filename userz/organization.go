package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.organization (
//	id varchar,
//	org_name varchar,
//
//	org_logo_bucket varchar,
//	org_logo_url varchar,
//
//	industry varchar,
//	org_type varchar,
//	zicops_subdomain varchar,
//	org_emp_count text,
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
		"org_name",
		"org_logo_bucket",
		"org_logo_url",
		"industry",
		"org_type",
		"zicops_subdomain",
		"org_emp_count",
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
	OrgName         string `db:"org_name"`
	OrgLogoBucket   string `db:"org_logo_bucket"`
	OrgLogoURL      string `db:"org_logo_url"`
	Industry        string `db:"industry"`
	OrgType         string `db:"org_type"`
	ZicopsSubdomain string `db:"zicops_subdomain"`
	OrgEmpCount     string `db:"org_emp_count"`
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
