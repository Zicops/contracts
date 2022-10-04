package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.resource (
//    type text,
//    url text,
//    topicId varchar,
//    created_at bigint,
//    updated_at bigint,
//    created_by text,
//    updated_by text,
//    PRIMARY KEY (topicId)
//);
// ResourceTableMeta is the table metadata for the resource table above.
var ResourceTableMeta = table.Metadata{
	Name: "resource",
	Columns: []string{
		"id",
		"name",
		"type",
		"bucketpath",
		"url",
		"topicid",
		"courseid",
		"lsp_id",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_active",
	},
	PartKey: []string{
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{},
}

// ResourceTable is the table for the resource table above.
var ResourceTable = table.New(ResourceTableMeta)

// define struct for resource table
type Resource struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Type       string `db:"type"`
	BucketPath string `db:"bucketpath"`
	Url        string `db:"url"`
	TopicId    string `db:"topicid"`
	CourseId   string `db:"courseid"`
	LspId      string `db:"lsp_id" json:"lsp_id"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
	CreatedBy  string `db:"created_by"`
	UpdatedBy  string `db:"updated_by"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}
