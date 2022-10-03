package qbankz

import "github.com/scylladb/gocqlx/table"

//create table exam_config(
//    id varchar,
//    exam_id varchar,
//    shuffle_questions boolean,
//    display_hints boolean,
//    show_answer boolean,
//    show_result boolean,
//    created_by varchar,
//    updated_by varchar,
//    created_at bigint,
//    updated_at bigint,
//    is_active boolean,
//
//    PRIMARY KEY (id)
//);
// ExamConfigTableMeta is the table metadata for the exam_config table above.
var ExamConfigTableMeta = table.Metadata{
	Name: "exam_config",
	Columns: []string{
		"id",
		"exam_id",
		"lsp_id",
		"shuffle_questions",
		"display_hints",
		"show_answer",
		"show_result",
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
	SortKey: []string{
		"created_at",
	},
}

// ExamConfigTable is the table for the exam_config table above.
var ExamConfigTable = table.New(ExamConfigTableMeta)

// define struct for exam_config table
type ExamConfig struct {
	ID           string `db:"id"`
	ExamID       string `db:"exam_id"`
	LspID        string `db:"lsp_id" json:"lsp_id"`
	Shuffle      bool   `db:"shuffle_questions"`
	DisplayHints bool   `db:"display_hints"`
	ShowAnswer   bool   `db:"show_answer"`
	ShowResult   bool   `db:"show_result"`
	CreatedBy    string `db:"created_by"`
	UpdatedBy    string `db:"updated_by"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
	IsActive     bool   `db:"is_active" json:"is_active"`
}
