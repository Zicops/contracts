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
		"quizId",
		"question",
		"correctAnswer",
		"explanation",
		"is_deleted",
	},
	PartKey: []string{
		"quizId",
	},
	SortKey: []string{},
}

// QuizDescriptiveTable is the table for the quiz_descriptive table above.
var QuizDescriptiveTable = table.New(QuizDescriptiveTableMeta)

// define struct for quiz_descriptive table
type QuizDescriptive struct {
	QuizId     string `db:"quizId"`
	Question   string `db:"question"`
	CorrectAnswer string `db:"correctAnswer"`
	Explanation string `db:"explanation"`
	IsDeleted  bool   `db:"is_deleted"`
}
