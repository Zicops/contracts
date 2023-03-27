package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.vendor_lsp_map (
//     vendor_id varchar,
//     lsp_id varchar,
//     is_source boolean,
//     services set<varchar>,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     type varchar,
//     words set<varchar>,
//     PRIMARY KEY((vendor_id, lsp_id), created_at)
// )
// WITH CLUSTERING ORDER BY (created_at DESC);

var VendorLspMapMeta = table.Metadata{
	Name: "vendor_lsp_map",
	Columns: []string{
		"vendor_id",
		"lsp_id",
		"is_source",
		"services",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
		"type",
		"words",
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
	VendorId  string   `db:"vendor_id"`
	LspId     string   `db:"lsp_id"`
	IsSource  bool     `db:"is_source"`
	Services  []string `db:"services"`
	CreatedAt int64    `db:"created_at"`
	CreatedBy string   `db:"created_by"`
	UpdatedAt int64    `db:"updated_at"`
	UpdatedBy string   `db:"updated_by"`
	Status    string   `db:"status"`
	Type      string   `db:"type"`
	Words     []string `db:"words"`
}
