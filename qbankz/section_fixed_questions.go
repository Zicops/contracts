package qbankz

import "github.com/scylladb/gocqlx/table"

// create table section_fixed_questions (
//     id varchar,
//     sqb_id varchar,
//     question_id varchar,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// SectionFixedQuestionsTableMeta is the table metadata for the section_fixed_questions table above.
var SectionFixedQuestionsTableMeta = table.Metadata{
	Name: "section_fixed_questions",
	Columns: []string{
		"id",
		"sqb_id",
		"lsp_id",
		"question_id",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"is_active",
	},
	SortKey: []string{
		"created_at",
	},
}

// SectionFixedQuestionsTable is the table for the section_fixed_questions table above.
var SectionFixedQuestionsTable = table.New(SectionFixedQuestionsTableMeta)

// define struct for section_fixed_questions table
type SectionFixedQuestions struct {
	ID         string `db:"id"`
	SQBId      string `db:"sqb_id"`
	LspID      string `db:"lsp_id"`
	QuestionID string `db:"question_id"`
	CreatedBy  string `db:"created_by"`
	UpdatedBy  string `db:"updated_by"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
	IsActive   bool   `db:"is_active"`
}
