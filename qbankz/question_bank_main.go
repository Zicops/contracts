package contracts

import "github.com/scylladb/gocqlx/table"

// create table qbankz.question_bank_main (
//     id varchar,
//     name text,
//     category text,
//     sub_category text,
//     owner varchar,
//     is_active boolean,
//     created_by varchar,
//     updated_by varchar,
//     created_at bigint,
//     updated_at bigint
//     is_default boolean,

//     PRIMARY KEY (id)
// );
// QuestionBankMainTableMeta is the table metadata for the question_bank_main table above.
var QuestionBankMainTableMeta = table.Metadata{
	Name: "question_bank_main",
	Columns: []string{
		"id",
		"name",
		"category",
		"sub_category",
		"owner",
		"is_active",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"is_default",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// QuestionBankMainTable is the table for the question_bank_main table above.
var QuestionBankMainTable = table.New(QuestionBankMainTableMeta)

// define struct for question_bank_main table
type QuestionBankMain struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Category    string `db:"category"`
	SubCategory string `db:"sub_category"`
	Owner       string `db:"owner"`
	IsActive    bool   `db:"is_active"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	IsDefault   bool   `db:"is_default"`
}
