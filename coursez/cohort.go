package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.course_cohort_mapping (
//    id varchar,
//    courseId varchar,
//    cohortId varchar,
//    course_type varchar,
//    lsp_id varchar,
//    cohort_code varchar,
//    is_mandatory boolean,
//    course_status varchar,
//    added_by text,
//    created_at bigint,
//    updated_at bigint,
//    created_by text,
//    updated_by text,
//    is_active boolean,
//    expected_completion_days int,
//    PRIMARY KEY (id)
//);
//
// CourseCohortMappingTableMeta is the table metadata for the course_cohort_mapping table above.
var CourseCohortMappingTableMeta = table.Metadata{
	Name: "course_cohort_mapping",
	Columns: []string{
		"id",
		"courseid",
		"cohortid",
		"course_type",
		"lsp_id",
		"cohort_code",
		"is_mandatory",
		"course_status",
		"added_by",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_active",
		"expected_completion_days",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{
		"created_at",
	},
}

var CourseCohortTable = table.New(CourseCohortMappingTableMeta)

// define struct for course_cohort_mapping table
type CourseCohortMapping struct {
	ID                     string `db:"id"`
	CourseID               string `db:"courseid"`
	CohortID               string `db:"cohortid"`
	CourseType             string `db:"course_type"`
	LspID                  string `db:"lsp_id" json:"lsp_id"`
	CohortCode             string `db:"cohort_code"`
	IsMandatory            bool   `db:"is_mandatory"`
	CourseStatus           string `db:"course_status"`
	AddedBy                string `db:"added_by"`
	CreatedAt              int64  `db:"created_at"`
	UpdatedAt              int64  `db:"updated_at"`
	CreatedBy              string `db:"created_by"`
	UpdatedBy              string `db:"updated_by"`
	IsActive               bool   `db:"is_active" json:"is_active"`
	ExpectedCompletionDays int    `db:"expected_completion_days"`
}
