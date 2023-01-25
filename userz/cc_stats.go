package userz

import "github.com/scylladb/gocqlx/table"

//
//create table userz.course_consumption_stats (
//	id varchar,
//	lsp_id varchar,
//	course_id varchar,
//	category varchar,
//	sub_category text,
//	owner varchar,
//	duration bigint,
//	total_learners bigint,
//	active_learners bigint,
//	completed_learners bigint,
//	expected_completion_time bigint,
//	average_completion_time bigint,
//	average_compliance_score bigint,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY ((lsp_id, course_id), total_learners)
//  )
//  WITH CLUSTERING ORDER BY (total_learners DESC);
//
var CCMain = table.Metadata{
	Name: "course_consumption_stats",
	Columns: []string{
		"id",
		"lsp_id",
		"course_id",
		"category",
		"sub_category",
		"owner",
		"duration",
		"total_learners",
		"active_learners",
		"completed_learners",
		"expected_completion_time",
		"average_completion_time",
		"average_compliance_score",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"lsp_id",
		"course_id",
	},
	SortKey: []string{
		"total_learners",
	},
}

var CCTable = table.New(CCMain)

type CCStats struct {
	ID                     string `db:"id"`
	LspId                  string `db:"lsp_id" json:"lsp_id"`
	CourseId               string `db:"course_id"`
	Category               string `db:"category"`
	SubCategory            string `db:"sub_category"`
	Owner                  string `db:"owner"`
	Duration               int64  `db:"duration"`
	TotalLearners          int64  `db:"total_learners"`
	ActiveLearners         int64  `db:"active_learners"`
	CompletedLearners      int64  `db:"completed_learners"`
	ExpectedCompletionTime int64  `db:"expected_completion_time"`
	AverageCompletionTime  int64  `db:"average_completion_time"`
	AverageComplianceScore int64  `db:"average_compliance_score"`
	CreatedAt              int64  `db:"created_at"`
	UpdatedAt              int64  `db:"updated_at"`
	CreatedBy              string `db:"created_by"`
	UpdatedBy              string `db:"updated_by"`
}
