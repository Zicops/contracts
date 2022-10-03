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
		"moduleid",
		"chapterid",
		"courseid",
		"lsp_id",
		"sequence",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"image",
		"imagebucket",
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

// TopicTable is the table for the topic table above.
var TopicTable = table.New(TopicTableMeta)

// define struct Topic for table coursez.topic
type Topic struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Type        string `db:"type"`
	ModuleID    string `db:"moduleid"`
	ChapterID   string `db:"chapterid"`
	CourseID    string `db:"courseid"`
	LspID       string `db:"lsp_id" json:"lsp_id"`
	Sequence    int    `db:"sequence"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
	Image       string `db:"image"`
	ImageBucket string `db:"imagebucket"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}
