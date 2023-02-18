package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// create table vendorz.sample_file (
//     sf_id varchar,
//     vendor_id varchar,
//     type varchar,
//     name varchar,
//     p_type varchar,
//     vendor_id varchar,
//     description text,
//     pricing_rate varchar,
//     file_type varchar,
//     file_url text,
//     file text,
//     file_bucket text,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((sf_id), created_at)
// )

var SampleFileMeta = table.Metadata{
	Name: "sample_file",
	Columns: []string{
		"sf_id",
		"name",
		"description",
		"pricing_rate",
		"file_type",
		"file_url",
		"file",
		"file_bucket",
		"p_type",
		"vendor_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"sf_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var SampleFileTable = table.New(SampleFileMeta)

type SampleFile struct {
	SfId        string `db:"sf_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Pricing     string `db:"pricing_rate"`
	FileType    string `db:"file_type"`
	FileUrl     string `db:"file_url"`
	File        string `db:"file"`
	FileBucket  string `db:"file_bucket"`
	PType       string `db:"p_type"`
	VendorId    string `db:"vendor_id"`
	CreatedAt   int64  `db:"created_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedAt   int64  `db:"updated_at"`
	UpdatedBy   string `db:"updated_by"`
	Status      string `db:"status"`
}
