package qbankz

import "github.com/scylladb/gocqlx/table"

// create table exam_schedule(
//     id varchar,
//     exam_id varchar,
//     start_date bigint,
//     end_date bigint,
//     buffer_time int,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// ExamScheduleTableMeta is the table metadata for the exam_schedule table above.
var ExamScheduleTableMeta = table.Metadata{
	Name: "exam_schedule",
	Columns: []string{
		"id",
		"exam_id",
		"lsp_id",
		"start",
		"end",
		"buffer_time",
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

// ExamScheduleTable is the table for the exam_schedule table above.
var ExamScheduleTable = table.New(ExamScheduleTableMeta)

// define struct for exam_schedule table
type ExamSchedule struct {
	ID         string `db:"id"`
	ExamID     string `db:"exam_id"`
	LspID      string `db:"lsp_id" json:"lsp_id"`
	Start      int64  `db:"start"`
	End        int64  `db:"end"`
	BufferTime int    `db:"buffer_time"`
	CreatedBy  string `db:"created_by"`
	UpdatedBy  string `db:"updated_by"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}
