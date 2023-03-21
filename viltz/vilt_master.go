package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.vilt_master (
// 	lsp_id varchar,
// 	course_id varchar,
// 	no_of_learners bigint,
// 	trainers set<varchar>,
// 	moderators set<varchar>,
// 	course_start_date bigint,
// 	course_end_date bigint,
// 	curriculum text,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
//  status varchar,
// 	primary key((course_id), created_at)
// 	)
// 	with clustering order by (created_at DESC);

var ViltMasterMeta = table.Metadata{
	Name: "vilt_master",
	Columns: []string{
		"lsp_id",
		"course_id",
		"no_of_learners",
		"trainers",
		"moderators",
		"course_start_date",
		"course_end_date",
		"curriculum",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"course_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var ViltMasterTable = table.New(ViltMasterMeta)

type ViltMaster struct {
	LspId           string   `db:"lsp_id"`
	CourseId        string   `db:"course_id"`
	NoOfLearners    int64    `db:"no_of_learners"`
	Trainers        []string `db:"trainers"`
	Moderators      []string `db:"moderators"`
	CourseStartDate int64    `db:"course_start_date"`
	CourseEndDate   int64    `db:"course_end_date"`
	Curriculum      string   `db:"curriculum"`
	CreatedAt       int64    `db:"created_at"`
	CreatedBy       string   `db:"created_by"`
	UpdatedAt       int64    `db:"updated_at"`
	UpdatedBy       string   `db:"updated_by"`
	Status          string   `db:"status"`
}
