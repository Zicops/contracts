package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// CREATE TABLE vendorz.vendor_order  (
// 	id varchar,
// 	vendor_id varchar,
//  lsp_id varchar,
// 	total bigint,
// 	tax bigint,
// 	grand_total bigint,
//  description varchar,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updated_by varchar,
// 	status varchar,
// 	PRIMARY KEY((vendor_id, id), created_at)
// );

var VendorOrderTableMeta = table.Metadata{
	Name: "vendor_order",
	Columns: []string{
		"id",
		"vendor_id",
		"lsp_id",
		"total",
		"tax",
		"grand_total",
		"description",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"id",
	},
	SortKey: []string{
		"created_at",
	},
}

var VendorOrderTable = table.New(VendorOrderTableMeta)

type VendorOrder struct {
	OrderId     string `db:"id"`
	VendorId    string `db:"vendor_id"`
	LspId       string `db:"lsp_id"`
	Total       int64  `db:"total"`
	Tax         int64  `db:"tax"`
	GrandTotal  int64  `db:"grand_total"`
	Description string `db:"description"`
	CreatedAt   int64  `db:"created_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedAt   int64  `db:"updated_at"`
	UpdatedBy   string `db:"updated_by"`
	Status      string `db:"status"`
}
