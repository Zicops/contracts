package qbankz

import "github.com/scylladb/gocqlx/table"

// create table exam(
//     id varchar,
//     name text,
//     description text,
//     qp_id varchar,
//     type varchar,
//     schedule_type varchar,
//     duratuoin int,
//     status varchar,
//     category text,
//     sub_category text,
//     is_active boolean,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// ExamTableMeta is the table metadata for the exam table above.
var ExamTableMeta = table.Metadata{
	Name: "exam",
	Columns: []string{
		"id",
		"lsp_id",
		"name",
		"words",
		"description",
		"code",
		"qp_id",
		"question_ids",
		"type",
		"schedule_type",
		"duration",
		"status",
		"category",
		"sub_category",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"is_active",
		"total_count",
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

// ExamTable is the table for the exam table above.
var ExamTable = table.New(ExamTableMeta)

// define struct for exam table
type Exam struct {
	ID           string   `db:"id"`
	LSPID        string   `db:"lsp_id" json:"lsp_id"`
	Name         string   `db:"name"`
	Words        []string `db:"words"`
	Description  string   `db:"description"`
	Code         string   `db:"code"`
	QPID         string   `db:"qp_id"`
	QuestionIDs  []string `db:"question_ids"`
	Type         string   `db:"type"`
	ScheduleType string   `db:"schedule_type"`
	Duration     int      `db:"duration"`
	Status       string   `db:"status"`
	Category     string   `db:"category"`
	SubCategory  string   `db:"sub_category"`
	IsActive     bool     `db:"is_active" json:"is_active"`
	CreatedBy    string   `db:"created_by"`
	UpdatedBy    string   `db:"updated_by"`
	CreatedAt    int64    `db:"created_at"`
	UpdatedAt    int64    `db:"updated_at"`
	TotalCount   int      `db:"total_count"`
}
