package vendorz

import "github.com/scylladb/gocqlx/table"

// create table vendorz.vendor_user_map (
//     vendor_id varchar,
//     user_id varchar,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((vendor_id, user_id), created_at)
// )

var VendorUserMapMeta = table.Metadata{
	Name: "vendor_user_map",
	Columns: []string{
		"vendor_id",
		"user_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"vendor_id",
		"user_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var VendorUserMapTable = table.New(VendorUserMapMeta)

type VendorUserMap struct {
	VendorId  string `db:"vendor_id"`
	UserId    string `db:"user_id"`
	CreatedAt int64  `db:"created_at"`
	CreatedBy string `db:"created_by"`
	UpdatedAt int64  `db:"updated_at"`
	UpdatedBy string `db:"updated_by"`
	Status    string `db:"status"`
}
