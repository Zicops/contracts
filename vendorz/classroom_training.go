package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.classroom_training (
//     ct_id varchar,
//     vendor_id varchar,
//     description text,
//     is_applicable boolean,
//     expertise set<varchar>,
//     languages set<varchar>,
//     output_deliveries set<varchar>,
//     sample_files set<varchar>,
//     is_expertise_online boolean,
//     is_expertise_offline boolean,
//     profiles set<varchar>,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, ct_id), created_at)
// )

var ClassRoomTrainingMeta = table.Metadata{
	Name: "classroom_training",
	Columns: []string{
		"ct_id",
		"vendor_id",
		"description",
		"is_applicable",
		"expertise",
		"languages",
		"output_deliveries",
		"sample_files",
		"is_expertise_online",
		"is_expertise_offline",
		"profiles",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"ct_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var ClassRoomTrainingTable = table.New(ClassRoomTrainingMeta)

// CRT - class room training
type CRT struct {
	CtId               string   `db:"ct_id"`
	VendorId           string   `db:"vendor_id"`
	Description        string   `db:"description"`
	IsApplicable       bool     `db:"is_applicable"`
	Expertise          []string `db:"expertise"`
	Languages          []string `db:"languages"`
	OutputDeliveries   []string `db:"output_deliveries"`
	SampleFiles        []string `db:"sample_files"`
	IsExpertiseOnline  bool     `db:"is_expertise_online"`
	IsExpertiseOffline bool     `db:"is_expertise_offline"`
	Profiles           []string `db:"profiles"`
	CreatedAt          int64    `db:"created_at"`
	CreatedBy          string   `db:"created_by"`
	UpdatedAt          int64    `db:"updated_at"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
}
