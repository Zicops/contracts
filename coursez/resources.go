package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.resource (
//    type text,
//    url text,
//    topicId varchar,
//    created_at timestamp,
//    updated_at timestamp,
//    created_by text,
//    updated_by text,
//    PRIMARY KEY (topicId)
//);
// ResourceTableMeta is the table metadata for the resource table above.
var ResourceTableMeta = table.Metadata{
	Name: "resource",
	Columns: []string{
		"type",
		"url",
		"topicId",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_deleted",
	},
	PartKey: []string{
		"topicId",
	},
	SortKey: []string{},
}

// ResourceTable is the table for the resource table above.
var ResourceTable = table.New(ResourceTableMeta)

// define struct for resource table
type Resource struct {
	Type       string `db:"type"`
	Url        string `db:"url"`
	TopicId    string `db:"topicId"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	CreatedBy  string `db:"created_by"`
	UpdatedBy  string `db:"updated_by"`
	IsDeleted  bool   `db:"is_deleted"`
}
