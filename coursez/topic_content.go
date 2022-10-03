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
		"id",
		"language",
		"topicid",
		"courseid",
		"lsp_id",
		"created_at",
		"updated_at",
		"starttime",
		"duration",
		"skipintroduration",
		"nextshowtime",
		"fromendtime",
		"url",
		"topiccontentbucket",
		"subtitlefile",
		"subtitlefilebucket",
		"is_active",
		"type",
		"is_default",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{
		"duration",
	},
}

// TopicContentTable is the table for the topic_content table above.
var TopicContentTable = table.New(TopicContentTableMeta)

// define struct for topic_content table
type TopicContent struct {
	ID                 string `db:"id"`
	Language           string `db:"language"`
	TopicId            string `db:"topicid"`
	CourseId           string `db:"courseid"`
	LspId              string `db:"lsp_id"`
	CreatedAt          int64  `db:"created_at"`
	UpdatedAt          int64  `db:"updated_at"`
	StartTime          int    `db:"starttime"`
	Duration           int    `db:"duration"`
	SkipIntroDuration  int    `db:"skipintroduration"`
	NextShowtime       int    `db:"nextshowtime"`
	FromEndTime        int    `db:"fromendtime"`
	TopicContentBucket string `db:"topiccontentbucket"`
	Url                string `db:"url"`
	SubtitleFile       string `db:"subtitlefile"`
	SubtitleFileBucket string `db:"subtitlefilebucket"`
	Type               string `db:"type"`
	IsActive           bool   `db:"is_active"`
	IsDefault          bool   `db:"is_default"`
}
