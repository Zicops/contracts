package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.organization_units(
//	id varchar,
//	org_id varchar,
//	org_unit_emp_count text,
//	address varchar,
//	country varchar,
//	state varchar,
//	city varchar,
//	postal_code text,
//	status varchar,
//	created_at bigint,
//	created_by varchar,
//	updated_at bigint,
//	updated_by varchar,
//	PRIMARY KEY ((id, org_id))
//  );

var OrganizationUnitsTableMeta = table.Metadata{
	Name: "organization_units",
	Columns: []string{
		"id",
		"org_id",
		"emp_count",
		"address",
		"country",
		"state",
		"city",
		"postal_code",
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
	SortKey: []string{},
}

var OrganizationUnitsTable = table.New(OrganizationUnitsTableMeta)

type OrganizationUnits struct {
	ID         string `db:"id"`
	OrgID      string `db:"org_id"`
	EmpCount   string `db:"emp_count"`
	Address    string `db:"address"`
	Country    string `db:"country"`
	State      string `db:"state"`
	City       string `db:"city"`
	PostalCode string `db:"postal_code"`
	Status     string `db:"status"`
	CreatedAt  int64  `db:"created_at"`
	CreatedBy  string `db:"created_by"`
	UpdatedAt  int64  `db:"updated_at"`
	UpdatedBy  string `db:"updated_by"`
}
