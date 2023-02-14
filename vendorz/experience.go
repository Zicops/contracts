package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// create table vendorz.experience (
//     exp_id varchar,
//     vendor_id varchar,
//     pf_id varchar,
//     start_date bigint,
//     end_date bigint,
//     title varchar,
//     location varchar,
//     location_type varchar,
//	   employement_type varchar,
//     company varchar,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, exp_id), created_at)
// )

var VendorExperienceMeta = table.Metadata{
	Name: "experience",
	Columns: []string{
		"exp_id",
		"vendor_id",
		"pf_id",
		"start_date",
		"end_date",
		"title",
		"location",
		"location_type",
		"employement_type",
		"company",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"exp_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var VendorExperienceTable = table.New(VendorExperienceMeta)

type VendorExperience struct {
	ExpId           string `db:"exp_id"`
	VendorId        string `db:"vendor_id"`
	PfId            string `db:"pf_id"`
	StartDate       int64  `db:"start_date"`
	EndDate         int64  `db:"end_date"`
	Title           string `db:"title"`
	Location        string `db:"location"`
	LocationType    string `db:"location_type"`
	EmployementType string `db:"employement_type"`
	Company         string `db:"company"`
	CreatedAt       int64  `db:"created_at"`
	CreatedBy       string `db:"created_by"`
	UpdatedAt       int64  `db:"updated_at"`
	UpdatedBy       string `db:"updated_by"`
	Status          string `db:"status"`
}
