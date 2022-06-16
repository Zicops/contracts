package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.topic_exam(
//    id varchar,
//    examId varchar,
//    topicId varchar,
//    language varchar,
//    created_at bigint,
//    updated_at bigint,
//    is_active boolean,
//    PRIMARY KEY (id)
//);

var TopicExamTableMeta = table.Metadata{
	Name: "topic_exam",
	Columns: []string{
		"id",
		"examId",
		"topicId",
		"language",
		"created_at",
		"updated_at",
		"is_active",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

type TopicExam struct {
	ID                 string `db:"id"`
	ExamId             string `db:"examId"`
	TopicId            string `db:"topicId"`
	Language           string `db:"language"`
	CreatedAt          int64  `db:"created_at"`
	UpdatedAt          int64  `db:"updated_at"`
	IsActive           bool   `db:"is_active"`
}
