package qbankz

import "github.com/scylladb/gocqlx/table"

// create table qbankz.question_main (
//     id varchar,
//     description text,
//     type varchar,
//     difficulty_score int,
//     attachment_url text,
//     attachment_type varchar,
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
		"description",
		"type",
		"difficulty_score",
		"attachment_url",
		"attachment_type",
		"qbm_id",
		"status",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// QuestionMainTable is the table for the question_main table above.
var QuestionMainTable = table.New(QuestionMainTableMeta)

// define struct for question_main table
type QuestionMain struct {
	ID             string `db:"id"`
	Description    string `db:"description"`
	Type           string `db:"type"`
	Difficulty     int    `db:"difficulty_score"`
	Attachment     string `db:"attachment_url"`
	AttachmentType string `db:"attachment_type"`
	QbmId          string `db:"qbm_id"`
	Status         string `db:"status"`
	CreatedBy      string `db:"created_by"`
	UpdatedBy      string `db:"updated_by"`
	CreatedAt      int64  `db:"created_at"`
	UpdatedAt      int64  `db:"updated_at"`
}
