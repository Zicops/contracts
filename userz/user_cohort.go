package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_cohort_map (
// 	id varchar,
// 	user_id varchar,
// 	user_lsp_id varchar,
// 	cohort_id varchar,
// 	added_by varchar,
// 	membership_status varchar,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserCohortMapTableMeta = table.Metadata{
	Name: "user_cohort_map",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"cohort_id",
		"added_by",
		"membership_status",
		"role",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"user_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserCohortTable = table.New(UserCohortMapTableMeta)

type UserCohort struct {
	ID               string `db:"id"`
	UserID           string `db:"user_id"`
	UserLspID        string `db:"user_lsp_id"`
	CohortID         string `db:"cohort_id"`
	AddedBy          string `db:"added_by"`
	MembershipStatus string `db:"membership_status"`
	Role             string `db:"role"`
	CreatedAt        int64  `db:"created_at"`
	UpdatedAt        int64  `db:"updated_at"`
	CreatedBy        string `db:"created_by"`
	UpdatedBy        string `db:"updated_by"`
}
