package qbankz

import "github.com/scylladb/gocqlx/table"

//create table exam_instructions(
//    id varchar,
//    exam_id varchar,
//    passing_criteria varchar,
//    no_attempts int,
//    access_type varchar,
//    updated_by varchar,
//    created_at bigint,
//    updated_at bigint,
//    is_active boolean,
//
//    PRIMARY KEY (id)
//);
// ExamInstructionsTableMeta is the table metadata for the exam_instructions table above.
var ExamInstructionsTableMeta = table.Metadata{
	Name: "exam_instructions",
	Columns: []string{
		"id",
		"exam_id",
		"lsp_id",
		"instructions",
		"passing_criteria",
		"no_attempts",
		"access_type",
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
	SortKey: []string{
		"created_at",
	},
}

// ExamInstructionsTable is the table for the exam_instructions table above.
var ExamInstructionsTable = table.New(ExamInstructionsTableMeta)

// define struct for exam_instructions table
type ExamInstructions struct {
	ID              string `db:"id"`
	ExamID          string `db:"exam_id"`
	LspId           string `db:"lsp_id" json:"lsp_id"`
	Instructions    string `db:"instructions"`
	PassingCriteria string `db:"passing_criteria"`
	NoAttempts      int    `db:"no_attempts"`
	AccessType      string `db:"access_type"`
	CreatedBy       string `db:"created_by"`
	UpdatedBy       string `db:"updated_by"`
	CreatedAt       int64  `db:"created_at"`
	UpdatedAt       int64  `db:"updated_at"`
	IsActive        bool   `db:"is_active" json:"is_active"`
}
