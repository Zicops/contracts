package qbankz

import "github.com/scylladb/gocqlx/table"

// create table qbankz.question_main (
//     id varchar,
//     description text,
//     type varchar,
//     difficulty_score int,
//     attachment_url text,
//     attachment_type varchar,
//	   attachment_bucket text,
//     hint text,
//     qbm_id varchar,
//     status varchar,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,

//     PRIMARY KEY (id)
// );
// QuestionMainTableMeta is the table metadata for the question_main table above.
var QuestionMainTableMeta = table.Metadata{
	Name: "question_main",
	Columns: []string{
		"id",
		"name",
		"description",
		"type",
		"difficulty_score",
		"attachment_url",
		"attachment_type",
		"attachment_bucket",
		"hint",
		"qbm_id",
		"lsp_id",
		"status",
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

// QuestionMainTable is the table for the question_main table above.
var QuestionMainTable = table.New(QuestionMainTableMeta)

// define struct for question_main table
type QuestionMain struct {
	ID               string `db:"id"`
	Name             string `db:"name"`
	Description      string `db:"description"`
	Type             string `db:"type"`
	Difficulty       int    `db:"difficulty_score"`
	Attachment       string `db:"attachment_url"`
	AttachmentType   string `db:"attachment_type"`
	AttachmentBucket string `db:"attachment_bucket"`
	Hint             string `db:"hint"`
	QbmId            string `db:"qbm_id"`
	LspId            string `db:"lsp_id" json:"lsp_id"`
	Status           string `db:"status"`
	CreatedBy        string `db:"created_by"`
	UpdatedBy        string `db:"updated_by"`
	CreatedAt        int64  `db:"created_at"`
	UpdatedAt        int64  `db:"updated_at"`
	IsActive         bool   `db:"is_active" json:"is_active"`
}
