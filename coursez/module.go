package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.module (
//    id varchar,
//    isChapter boolean,
//    name text,
//    description text,
//    courseId varchar,
//    owner text,
//    created_at bigint,
//    updated_at bigint,
//    level text,
//    sequence int,
//    setGlobal boolean,
//    PRIMARY KEY ((courseId), id)
//) WITH CLUSTERING
//ORDER BY
//    (id ASC);
// ModuleTableMeta is the table metadata for the module table above.
var ModuleTableMeta = table.Metadata{
	Name: "module",
	Columns: []string{
		"id",
		"isChapter",
		"name",
		"description",
		"courseId",
		"owner",
		createdAt,
		updatedAt,
		"level",
		"sequence",
		"setGlobal",
		isDeleted,
	},
	PartKey: []string{
		"courseId",
		"id",
	},
	SortKey: []string{
		"id",
	},
}

// ModuleTable is the table for the module table above.
var ModuleTable = table.New(ModuleTableMeta)

// define struct Module for table coursez.module
type Module struct {
	ID          string `db:"id"`
	IsChapter   bool   `db:"isChapter"`
	Name        string `db:"name"`
	Description string `db:"description"`
	CourseID    string `db:"courseId"`
	Owner       string `db:"owner"`
	CreatedAt   int64  `db:createdAt`
	UpdatedAt   int64  `db:updatedAt`
	Level       string `db:"level"`
	Sequence    int    `db:"sequence"`
	SetGlobal   bool   `db:"setGlobal"`
	IsDeleted   bool   `db:isDeleted`
}
