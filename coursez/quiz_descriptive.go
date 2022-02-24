package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.quiz_descriptive(
//    quizId varchar,
//    question text,
//    correctAnswer text,
//    explanation text,
//    PRIMARY KEY (quizId)
//);
// QuizDescriptiveTableMeta is the table metadata for the quiz_descriptive table above.
var QuizDescriptiveTableMeta = table.Metadata{
	Name: "quiz_descriptive",
	Columns: []string{
		"quizid",
		"question",
		"correctanswer",
		"explanation",
		"is_active",
	},
	PartKey: []string{
		"quizid",
	},
	SortKey: []string{},
}

// QuizDescriptiveTable is the table for the quiz_descriptive table above.
var QuizDescriptiveTable = table.New(QuizDescriptiveTableMeta)

// define struct for quiz_descriptive table
type QuizDescriptive struct {
	QuizId        string `db:"quizid"`
	Question      string `db:"question"`
	CorrectAnswer string `db:"correctanswer"`
	Explanation   string `db:"explanation"`
	IsActive      bool   `db:"is_active"`
}
