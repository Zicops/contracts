package qbankz

import "github.com/scylladb/gocqlx/table"

// create table section_main (
//     id varchar,
//     name text,
//     description text,
//     qp_id varchar,
//     type varchar,
//     difficulty_level varchar,
//     total_questions int,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// SectionMainTableMeta is the table metadata for the section_main table above.
var SectionMainTableMeta = table.Metadata{
	Name: "section_main",
	Columns: []string{
		"id",
		"name",
		"description",
		"qp_id",
		"type",
		"difficulty_level",
		"total_questions",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"is_active",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// SectionMainTable is the table for the section_main table above.
var SectionMainTable = table.New(SectionMainTableMeta)

// define struct for section_main table
type SectionMain struct {
	ID              string `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	QPID            string `db:"qp_id"`
	Type            string `db:"type"`
	DifficultyLevel string `db:"difficulty_level"`
	TotalQuestions  int    `db:"total_questions"`
	CreatedBy       string `db:"created_by"`
	UpdatedBy       string `db:"updated_by"`
	CreatedAt       int64  `db:"created_at"`
	UpdatedAt       int64  `db:"updated_at"`
	IsActive        bool   `db:"is_active"`
}
