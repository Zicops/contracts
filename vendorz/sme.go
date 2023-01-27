package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.sme (
//     sme_id varchar,
//     vendor_id varchar,
//     description text,
//     is_applicable boolean,
//     expertise set<varchar>,
//     languages set<varchar>,
//     output_deliveries set<varchar>,
//     sample_files set<varchar>.
//     profiles set<varchar>,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, sme_id), created_at)
// )

var SMEMeta = table.Metadata{
	Name: "sme",
	Columns: []string{
		"sme_id",
		"vendor_id",
		"description",
		"is_applicable",
		"expertise",
		"languages",
		"output_deliveries",
		"sample_files",
		"profiles",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"sme_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var SMETable = table.New(SMEMeta)

type SME struct {
	SMEId            string   `db:"sme_id"`
	VendorId         string   `db:"vendor_id"`
	Description      string   `db:"descrption"`
	IsApplicable     bool     `db:"is_applicable"`
	Expertise        []string `db:"expertise"`
	Languages        []string `db:"languages"`
	OutputDeliveries []string `db:"output_deliveries"`
	SampleFiles      []string `db:"sample_files"`
	Profiles         []string `db:"profiles"`
	CreatedAt        int64    `db:"created_at"`
	CreatedBy        string   `db:"created_by"`
	UpdatedAt        int64    `db:"updated_at"`
	UpdatedBy        string   `db:"updated_by"`
	Status           string   `db:"status"`
}
