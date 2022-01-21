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
	},
	PartKey: []string{
		"topicId",
	},
	SortKey: []string{},
}

// TopicContentTable is the table for the topic_content table above.
var TopicContentTable = table.New(TopicContentTableMeta)
