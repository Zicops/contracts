package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// create table vendorz.order_services (
// 	id varchar,
// 	order_id varchar,
// 	service_type varchar,
// 	description varchar,
// 	unit varchar,
// 	currency varchar,
// 	rate bigint,
// 	quantity bigint,
// 	total bigint,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updated_by varchar,
// 	status varchar,
// 	PRIMARY KEY((id))
// );

var OrderServiesTableMeta = table.Metadata{
	Name: "order_services",
	Columns: []string{
		"id",
		"order_id",
		"service_type",
		"description",
		"unit",
		"currency",
		"rate",
		"quantity",
		"total",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"id",
	},
}

var OrderServiesTable = table.New(OrderServiesTableMeta)

type OrderServices struct {
	ServiceId   string `db:"id"`
	OrderId     string `db:"order_id"`
	ServiceType string `db:"service_type"`
	Description string `db:"description"`
	Unit        string `db:"unit"`
	Currency    string `db:"currency"`
	Rate        int64  `db:"rate"`
	Quantity    int64  `db:"quantity"`
	Total       int64  `db:"total"`
	CreatedAt   int64  `db:"created_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedAt   int64  `db:"updated_at"`
	UpdatedBy   string `db:"updated_by"`
	Status      string `db:"status"`
}
