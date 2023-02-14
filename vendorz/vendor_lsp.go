package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.vendor_lsp_map (
//     vendor_id varchar,
//     lsp_id varchar,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, lsp_id), created_at)
// )
// WITH CLUSTERING ORDER BY (created_at DESC);

var VendorLspMapMeta = table.Metadata{
	Name: "vendor_lsp_map",
	Columns: []string{
		"vendor_id",
		"lsp_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"lsp_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var VendorLspMapTable = table.New(VendorLspMapMeta)

type VendorLspMap struct {
	VendorId  string `db:"vendor_id"`
	LspId     string `db:"lsp_id"`
	CreatedAt int64  `db:"created_at"`
	CreatedBy string `db:"created_by"`
	UpdatedAt int64  `db:"updated_at"`
	UpdatedBy string `db:"updated_by"`
	Status    string `db:"status"`
}