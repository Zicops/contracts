package vendorz

import (
	"github.com/scylladb/gocqlx/table"
)

// create table vendorz.vendor (
//     id varchar,
//     type varchar,
//     level varchar,
//     lsp_id varchar,
//     phone varchar,
//     name text,
//     photo text,
//     photo_url varchar,
//     photo_bucket varchar,
//     address text,
//     website varchar,
//     facebook varchar,
//     instagram varchar,
//     twitter varchar,
//     linkedin varchar,
//     users set<varchar>,
//     description text,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((id), created_at)
// )

var VendorTableMeta = table.Metadata{
	Name: "vendor",
	Columns: []string{
		"id",
		"type",
		"level",
		"lsp_id",
		"phone",
		"name",
		"photo",
		"photo_url",
		"photo_bucket",
		"address",
		"website",
		"facebook",
		"instagram",
		"twitter",
		"linkedin",
		"users",
		"description",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{
		"created_at",
	},
}

var VendorTable = table.New(VendorTableMeta)

type Vendor struct {
	VendorId    string   `db:"id"`
	Type        string   `db:"type"`
	Level       string   `db:"level"`
	LspId       string   `db:"lsp_id"`
	Phone       string   `db:"phone"`
	Name        string   `db:"name"`
	Photo       string   `db:"photo"`
	PhotoUrl    string   `db:"photo_url"`
	PhotoBucket string   `db:"photo_bucket"`
	Address     string   `db:"address"`
	Website     string   `db:"website"`
	Facebook    string   `db:"facebook"`
	Instagram   string   `db:"instagram"`
	Twitter     string   `db:"twitter"`
	LinkedIn    string   `db:"linkedin"`
	Users       []string `db:"users"`
	Description string   `db:"description"`
	CreatedAt   int64    `db:"created_at"`
	CreatedBy   string   `db:"created_by"`
	UpdatedAt   int64    `db:"updated_at"`
	UpdatedBy   string   `db:"updated_by"`
	Status      string   `db:"status"`
}
