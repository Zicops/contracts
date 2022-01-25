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
		"type",
		"url",
		"topicId",
		createdAt,
		updatedAt,
		createdBy,
		updatedBy,
		isDeleted,
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
	BucketPath string `db:"bucketPath"`
	Url        string `db:"url"`
	TopicId    string `db:"topicId"`
	CreatedAt  int64  `db:createdAt`
	UpdatedAt  int64  `db:updatedAt`
	CreatedBy  string `db:createdBy`
	UpdatedBy  string `db:updatedBy`
	IsDeleted  bool   `db:isDeleted`
}
