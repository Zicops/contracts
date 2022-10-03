package qbankz

import "github.com/scylladb/gocqlx/table"

// create table question_paper_main (
//     id varchar,
//     name text,
//     description text,
//     category text,
//     sub_category text,
//     is_active boolean,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,
//	   status varchar,
//     PRIMARY KEY (id)
// );
// QuestionPaperMainTableMeta is the table metadata for the question_paper_main table above.
var QuestionPaperMainTableMeta = table.Metadata{
	Name: "question_paper_main",
	Columns: []string{
		"id",
		"lsp_id",
		"name",
		"description",
		"section_wise",
		"difficulty_level",
		"suggested_duration",
		"category",
		"sub_category",
		"is_active",
		"status",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
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

// QuestionPaperMainTable is the table for the question_paper_main table above.
var QuestionPaperMainTable = table.New(QuestionPaperMainTableMeta)

// define struct for question_paper_main table
type QuestionPaperMain struct {
	ID                string `db:"id"`
	LspID             string `db:"lsp_id"`
	Name              string `db:"name"`
	Description       string `db:"description"`
	SectionWise       bool   `db:"section_wise"`
	DifficultyLevel   string `db:"difficulty_level"`
	SuggestedDuration string `db:"suggested_duration"`
	Category          string `db:"category"`
	SubCategory       string `db:"sub_category"`
	IsActive          bool   `db:"is_active"`
	Status            string `db:"status"`
	CreatedBy         string `db:"created_by"`
	UpdatedBy         string `db:"updated_by"`
	CreatedAt         int64  `db:"created_at"`
	UpdatedAt         int64  `db:"updated_at"`
}
