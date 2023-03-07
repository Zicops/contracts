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
}
