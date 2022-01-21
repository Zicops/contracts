package coursez

import "github.com/scylladb/gocqlx/v2/table"

// create table coursez.topic (
//    id varchar,
//    name text,
//    description text,
//    type varchar,
//    moduleId varchar,
//    chapterId varchar,
//    courseId varchar,
//    sequence int,
//    created_at timestamp,
//    updated_at timestamp,
//    created_by text,
//    updated_by text,
//    PRIMARY KEY ((courseId, moduleId), id)
//) WITH CLUSTERING
//ORDER BY
//    (id ASC);
// TopicTableMeta is the table metadata for the topic table above.
var TopicTableMeta = table.Metadata{
	Name: "topic",
	Columns: []string{
		"id",
		"name",
		"description",
		"type",
		"moduleId",
		"chapterId",
		"courseId",
		"sequence",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"courseId",
		"moduleId",
		"id",
	},
	SortKey: []string{
		"id",
	},
}

// TopicTable is the table for the topic table above.
var TopicTable = table.New(TopicTableMeta)
