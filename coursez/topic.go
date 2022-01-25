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
//    created_at bigint,
//    updated_at bigint,
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
		createdAt,
		updatedAt,
		createdBy,
		updatedBy,
		isDeleted,
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

// define struct Topic for table coursez.topic
type Topic struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Type        string `db:"type"`
	ModuleID    string `db:"moduleId"`
	ChapterID   string `db:"chapterId"`
	CourseID    string `db:"courseId"`
	Sequence    int    `db:"sequence"`
	CreatedAt   int64  `db:createdAt`
	UpdatedAt   int64  `db:updatedAt`
	CreatedBy   string `db:createdBy`
	UpdatedBy   string `db:updatedBy`
	IsDeleted   bool   `db:isDeleted`
}
