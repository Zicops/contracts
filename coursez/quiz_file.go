package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.quiz_file(
//    quizId varchar,
//    type text,
//    name text,
//    path text,
//    PRIMARY KEY (quizId)
//);
// QuizFileTableMeta is the table metadata for the quiz_file table above.
var QuizFileTableMeta = table.Metadata{
	Name: "quiz_file",
	Columns: []string{
		"id",
		"quizid",
		"lsp_id",
		"type",
		"name",
		"bucketpath",
		"path",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{},
}

// QuizFileTable is the table for the quiz_file table above.
var QuizFileTable = table.New(QuizFileTableMeta)

// define struct for quiz_file table
type QuizFile struct {
	ID         string `db:"id"`
	QuizId     string `db:"quizid"`
	LspId      string `db:"lsp_id" json:"lsp_id"`
	Type       string `db:"type"`
	Name       string `db:"name"`
	BucketPath string `db:"bucketpath"`
	Path       string `db:"path"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}
