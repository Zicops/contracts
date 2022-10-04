package qbankz

import "github.com/scylladb/gocqlx/table"

//create table exam_cohort(
//    id varchar,
//    exam_id varchar,
//    cohort_id varchar,
//    created_by varchar,
//    updated_by varchar,
//    created_at bigint,
//    updated_at bigint,
//    is_active boolean,
//
//    PRIMARY KEY (id)
//);
// ExamCohortTableMeta is the table metadata for the exam_cohort table above.
var ExamCohortTableMeta = table.Metadata{
	Name: "exam_cohort",
	Columns: []string{
		"id",
		"exam_id",
		"lsp_id",
		"cohort_id",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{},
}

// ExamCohortTable is the table for the exam_cohort table above.
var ExamCohortTable = table.New(ExamCohortTableMeta)

// define struct for exam_cohort table
type ExamCohort struct {
	ID        string `db:"id"`
	ExamID    string `db:"exam_id"`
	LspId     string `db:"lsp_id" json:"lsp_id"`
	CohortID  string `db:"cohort_id"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	IsActive  bool   `db:"is_active" json:"is_active"`
}
