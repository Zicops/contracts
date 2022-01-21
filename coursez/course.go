package coursez

import (
	"github.com/scylladb/gocqlx/v2/table"
)

//create table coursez.course (
//    id varchar,
//    name text,
//    description text,
//    instructor text,
//    image text,
//    previewVideo text,
//    owner text,
//    duration int,
//    level text,
//    language
//    set
//        < varchar >,
//        takeAways
//    set
//        < varchar >,
//        created_at timestamp,
//        updated_at timestamp,
//        type text,
//        prequisites
//    set
//        < varchar >,
//        goodFor
//    set
//        < varchar >,
//        mustFor
//    set
//        < varchar >,
//        created_by text,
//        updated_by text,
//        PRIMARY KEY ((id), created_at, type, level)
//) WITH CLUSTERING
//ORDER BY
//    (created_at DESC, type ASC, level ASC);
// CourseTableMeta is the table metadata for the course table above.
var CourseTableMeta = table.Metadata{
	Name: "course",
	Columns: []string{
		"id",
		"name",
		"description",
		"instructor",
		"image",
		"previewVideo",
		"owner",
		"duration",
		"level",
		"language",
		"set",
		"takeAways",
		"set",
		"created_at",
		"updated_at",
		"type",
		"prequisites",
		"set",
		"goodFor",
		"set",
		"mustFor",
		"set",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"id",
		"created_at",
		"type",
		"level",
	},
	SortKey: []string{
		"created_at",
		"type",
		"level",
	},
}

// CourseTable is the table for the course table above.
var CourseTable = table.New(CourseTableMeta)