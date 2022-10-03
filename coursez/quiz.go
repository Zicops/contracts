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
//    courseId varchar,
//    questionId varchar,
//	  qb_id varchar,
//    weightage int,
//    startTime int,
//    sequence int,
//    is_active boolean,
//    PRIMARY KEY (id)
//);

// QuizTableMeta is the table metadata for the quiz table above.
var QuizTableMeta = table.Metadata{
	Name: "quiz",
	Columns: []string{
		"id",
		"name",
		"category",
		"type",
		"ismandatory",
		"created_at",
		"updated_at",
		"topicid",
		"courseid",
		"lsp_id",
		"questionid",
		"weightage",
		"qbid",
		"starttime",
		"sequence",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{
		"sequence",
	},
}

// QuizTable is the table for the quiz table above.
var QuizTable = table.New(QuizTableMeta)

// define struct for quiz table
type Quiz struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Category    string `db:"category"`
	Type        string `db:"type"`
	IsMandatory bool   `db:"ismandatory"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	TopicID     string `db:"topicid"`
	CourseID    string `db:"courseid"`
	LspID       string `db:"lsp_id" json:"lsp_id"`
	QuestionID  string `db:"questionid"`
	QbId        string `db:"qbid"`
	Weightage   int    `db:"weightage"`
	StartTime   int    `db:"starttime"`
	Sequence    int    `db:"sequence"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}
