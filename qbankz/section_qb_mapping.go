package contracts

import "github.com/scylladb/gocqlx/table"

// create table section_qb_mapping (
//     id varchar,
//     qb_id varchar,
//     section_id varchar,
//     difficulty_level varchar,
//     total_questions int,
//     question_type varchar,
//     retrieval_type varchar,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// SectionQBMappingTableMeta is the table metadata for the section_qb_mapping table above.
var SectionQBMappingTableMeta = table.Metadata{
	Name: "section_qb_mapping",
	Columns: []string{
		"id",
		"qb_id",
		"section_id",
		"difficulty_level",
		"total_questions",
		"question_type",
		"retrieval_type",
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

// SectionQBMappingTable is the table for the section_qb_mapping table above.
var SectionQBMappingTable = table.New(SectionQBMappingTableMeta)

// define struct for section_qb_mapping table
type SectionQBMapping struct {
	ID              string `db:"id"`
	QBId            string `db:"qb_id"`
	SectionID       string `db:"section_id"`
	DifficultyLevel string `db:"difficulty_level"`
	TotalQuestions  int    `db:"total_questions"`
	QuestionType    string `db:"question_type"`
	RetrievalType   string `db:"retrieval_type"`
	CreatedBy       string `db:"created_by"`
	UpdatedBy       string `db:"updated_by"`
	CreatedAt       int64  `db:"created_at"`
	UpdatedAt       int64  `db:"updated_at"`
	IsActive        bool   `db:"is_active"`
}
