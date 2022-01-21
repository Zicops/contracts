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
		"quizId",
		"question",
		"options",
		"correctOption",
		"explanation",
	},
	PartKey: []string{
		"quizId",
	},
	SortKey: []string{},
}

// QuizMcqTable is the table for the quiz_mcq table above.
var QuizMcqTable = table.New(QuizMcqTableMeta)
