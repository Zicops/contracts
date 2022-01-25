package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.topic_content (
//    language
//    set
//        < text >,
//        topicId varchar,
//        created_at bigint,
//        updated_at bigint,
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
		"topicId",
		createdAt,
		updatedAt,
		"startTime",
		"duration",
		"skipIntro",
		"nextShowtime",
		"fromEndTime",
		"url",
		isDeleted,
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
	Language           string `db:"language"`
	TopicId            string `db:"topicId"`
	CreatedAt          int64  `db:createdAt`
	UpdatedAt          int64  `db:updatedAt`
	StartTime          int    `db:"startTime"`
	Duration           int    `db:"duration"`
	SkipIntro          bool   `db:"skipIntro"`
	NextShowtime       int    `db:"nextShowtime"`
	FromEndTime        int    `db:"fromEndTime"`
	TopicContentBucket string `db:"topicContentBucket"`
	Url                string `db:"url"`
	IsDeleted          bool   `db:isDeleted`
}
