package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.trainer(
// 	id varchar,
// 	lsp_id varchar,
// 	user_id varchar,
// 	vendor_id varchar,
// 	expertise set<varchar>,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updated_by bigint,
// 	status varchar,
// 	PRIMARY KEY((id))
// 	);

var ViltTrainerMeta = table.Metadata{
	Name: "trainer",
	Columns: []string{
		"id",
		"lsp_id",
		"user_id",
		"vendor_id",
		"expertise",
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

var ViltTrainerTable = table.New(ViltTrainerMeta)

type ViltTrainer struct {
	TrainerId string   `db:"id"`
	LspId     string   `db:"lsp_id"`
	UserId    string   `db:"user_id"`
	VendorId  string   `db:"vendor_id"`
	Expertise []string `db:"expertise"`
	CreatedAt int64    `db:"created_at"`
	CreatedBy string   `db:"created_by"`
	UpdatedAt int64    `db:"updated_at"`
	UpdatedBy string   `db:"updated_by"`
	status    string   `db:"status"`
}
