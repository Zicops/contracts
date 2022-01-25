package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.quiz (
//    id varchar,
//    name text,
//    category text,
//    type text,
//    isMandatory boolean,
//    created_at bigint,
//    updated_at bigint,
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
		createdAt,
		updatedAt,
		"topicId",
		"startTime",
		"sequence",
		isDeleted,
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
	CreatedAt   int64  `db:createdAt`
	UpdatedAt   int64  `db:updatedAt`
	TopicID     string `db:"topicId"`
	StartTime   int    `db:"startTime"`
	Sequence    int    `db:"sequence"`
	IsDeleted   bool   `db:isDeleted`
}
