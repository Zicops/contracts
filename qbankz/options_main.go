package contracts

import "github.com/scylladb/gocqlx/table"

// create table options_main (
//     id varchar,
//     description text,
//     is_correct boolean,
//     qm_id varchar,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint,
//     attachment_type varchar,
//     attachment_url text,
//     attachment_bucket varchar,
//     is_active boolean,

//     PRIMARY KEY (id)
// );
// OptionsMainTableMeta is the table metadata for the options_main table above.
var OptionsMainTableMeta = table.Metadata{
	Name: "options_main",
	Columns: []string{
		"id",
		"description",
		"is_correct",
		"qm_id",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"attachment_type",
		"attachment_url",
		"attachment_bucket",
		"is_active",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// OptionsMainTable is the table for the options_main table above.
var OptionsMainTable = table.New(OptionsMainTableMeta)

// define struct for options_main table
type OptionsMain struct {
	ID               string `db:"id"`
	Description      string `db:"description"`
	IsCorrect        bool   `db:"is_correct"`
	QmId             string `db:"qm_id"`
	CreatedBy        string `db:"created_by"`
	UpdatedBy        string `db:"updated_by"`
	CreatedAt        int64  `db:"created_at"`
	UpdatedAt        int64  `db:"updated_at"`
	AttachmentType   string `db:"attachment_type"`
	Attachment       string `db:"attachment_url"`
	AttachmentBucket string `db:"attachment_bucket"`
	IsActive         bool   `db:"is_active"`
}
