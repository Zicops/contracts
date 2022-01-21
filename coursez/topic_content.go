package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.topic_content (
//    language
//    set
//        < text >,
//        topicId varchar,
//        created_at timestamp,
//        updated_at timestamp,
//        startTime int,
//        duration int,
//        skipIntro boolean,
//        nextShowtime int,
//        fromEndTime int,
//        url text,
//        PRIMARY KEY (topicId)
//);
// TopicContentTableMeta is the table metadata for the topic_content table above.
var TopicContentTableMeta = table.Metadata{
	Name: "topic_content",
	Columns: []string{
		"language",
		"set",
		"topicId",
		"created_at",
		"updated_at",
		"startTime",
		"duration",
		"skipIntro",
		"nextShowtime",
		"fromEndTime",
		"url",
		"is_deleted",
	},
	PartKey: []string{
		"topicId",
	},
	SortKey: []string{},
}

// TopicContentTable is the table for the topic_content table above.
var TopicContentTable = table.New(TopicContentTableMeta)

// define struct for topic_content table
type TopicContent struct {
	Language   string `db:"language"`
	Set        string `db:"set"`
	TopicId    string `db:"topicId"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	StartTime  int    `db:"startTime"`
	Duration   int    `db:"duration"`
	SkipIntro  bool   `db:"skipIntro"`
	NextShowtime int    `db:"nextShowtime"`
	FromEndTime int    `db:"fromEndTime"`
	Url        string `db:"url"`
	IsDeleted  bool   `db:"is_deleted"`
}
