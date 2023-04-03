package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.content_development (
//     cd_id varchar,
//     vendor_id varchar,
//     description text,
//     is_applicable boolean,
//     expertise set<varchar>,
//     languages set<varchar>,
//     output_deliveries set<varchar>,
//     sample_files set<varchar>,
//     profiles set<varchar>,
//     is_expertise_online boolean,
//     is_expertise_offline boolean,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, cd_id), created_at)
// )

var ContentDevelopmentMeta = table.Metadata{
	Name: "content_development",
	Columns: []string{
		"cd_id",
		"vendor_id",
		"description",
		"is_applicable",
		"expertise",
		"languages",
		"output_deliveries",
		"sample_files",
		"profiles",
		"is_expertise_online",
		"is_expertise_offline",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"cd_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var ContentDevelopmentTable = table.New(ContentDevelopmentMeta)

type ContentDevelopment struct {
	CdId               string   `db:"cd_id"`
	VendorId           string   `db:"vendor_id"`
	Description        string   `db:"description"`
	IsApplicable       bool     `db:"is_applicable"`
	Expertise          []string `db:"expertise"`
	Languages          []string `db:"languages"`
	OutputDeliveries   []string `db:"output_deliveries"`
	SampleFiles        []string `db:"sample_files"`
	Profiles           []string `db:"profiles"`
	IsExpertiseOffline bool     `db:"is_expertise_offline"`
	IsExpertiseOnline  bool     `db:"is_expertise_online"`
	CreatedAt          int64    `db:"created_at"`
	CreatedBy          string   `db:"created_by"`
	UpdatedAt          int64    `db:"updated_at"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
}
