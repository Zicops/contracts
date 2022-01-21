package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.quiz (
//    id varchar,
//    name text,
//    category text,
//    type text,
//    isMandatory boolean,
//    created_at timestamp,
//    updated_at timestamp,
//    topicId varchar,
//    startTime int,
//    sequence int,
//    PRIMARY KEY ((topicId), id)
//);
// QuizTableMeta is the table metadata for the quiz table above.
var QuizTableMeta = table.Metadata{
	Name: "quiz",
	Columns: []string{
		"id",
		"name",
		"category",
		"type",
		"isMandatory",
		"created_at",
		"updated_at",
		"topicId",
		"startTime",
		"sequence",
		"is_deleted",
	},
	PartKey: []string{
		"topicId",
		"id",
	},
	SortKey: []string{},
}

// QuizTable is the table for the quiz table above.
var QuizTable = table.New(QuizTableMeta)

// define struct for quiz table
type Quiz struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Category    string `db:"category"`
	Type        string `db:"type"`
	IsMandatory bool   `db:"isMandatory"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	TopicID     string `db:"topicId"`
	StartTime   int    `db:"startTime"`
	Sequence    int    `db:"sequence"`
	IsDeleted   bool   `db:"is_deleted"`
}
