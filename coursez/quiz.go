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
	},
	PartKey: []string{
		"topicId",
		"id",
	},
	SortKey: []string{},
}

// QuizTable is the table for the quiz table above.
var QuizTable = table.New(QuizTableMeta)
