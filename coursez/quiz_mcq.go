package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.quiz_mcq(
//    quizId varchar,
//    question text,
//    options list<text>,
//    correctOption text,
//    explanation text,
//    PRIMARY KEY (quizId)
//);
// QuizMcqTableMeta is the table metadata for the quiz_mcq table above.
var QuizMcqTableMeta = table.Metadata{
	Name: "quiz_mcq",
	Columns: []string{
		"id",
		"quizid",
		"lsp_id",
		"question",
		"options",
		"correctoption",
		"explanation",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{
		"question",
	},
}

// QuizMcqTable is the table for the quiz_mcq table above.
var QuizMcqTable = table.New(QuizMcqTableMeta)

// define struct for quiz_mcq table
type QuizMcq struct {
	ID            string   `db:"id"`
	QuizId        string   `db:"quizid"`
	LspId         string   `db:"lsp_id" json:"lsp_id"`
	Question      string   `db:"question"`
	Options       []string `db:"options"`
	CorrectOption string   `db:"correctoption"`
	Explanation   string   `db:"explanation"`
	IsActive      bool     `db:"is_active" json:"is_active"`
}
