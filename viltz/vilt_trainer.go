package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.trainer(
// 	id varchar,
// 	lsp_id varchar,
// 	user_id varchar,
// 	vendor_id varchar,
// 	expertise set<varchar>,
//  website varchar,
//  github varchar,
//  years_of_experience varchar,
//  linkedin varchar,
//  description varchar,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updatedby varchar,
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
		"years_of_experience",
		"website",
		"github",
		"linkedin",
		"description",
		"created_at",
		"created_by",
		"updated_at",
		"updatedby",
		"status",
	},
	PartKey: []string{
		"id",
	},
}

var ViltTrainerTable = table.New(ViltTrainerMeta)

type ViltTrainer struct {
	TrainerId         string   `db:"id"`
	LspId             string   `db:"lsp_id"`
	UserId            string   `db:"user_id"`
	VendorId          string   `db:"vendor_id"`
	Expertise         []string `db:"expertise"`
	YearsOfExperience string   `db:"years_of_experience"`
	Website           string   `db:"website"`
	Github            string   `db:"github"`
	LinkedIn          string   `db:"linkedin"`
	Description       string   `db:"description"`
	CreatedAt         int64    `db:"created_at"`
	CreatedBy         string   `db:"created_by"`
	UpdatedAt         int64    `db:"updated_at"`
	UpdatedBy         string   `db:"updatedby"`
	Status            string   `db:"status"`
}
