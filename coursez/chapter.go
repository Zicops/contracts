package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.chapter (
//    id uuid,
//    name text,
//    description text,
//    moduleId uuid,
//    courseId uuid,
//    sequence int,
//    created_at bigint,
//    updated_at bigint,
//    PRIMARY KEY ((courseId), id)
//) WITH CLUSTERING
//ORDER BY
//    (id ASC);
// ChapterTableMeta is the table metadata for the chapter table above.
var ChapterTableMeta = table.Metadata{
	Name: "chapter",
	Columns: []string{
		"id",
		"name",
		"description",
		"moduleId",
		"courseId",
		"sequence",
		"created_at",
		"updated_at",
		"is_active",
	},
	PartKey: []string{
		"courseId",
		"id",
	},
	SortKey: []string{
		"id",
	},
}

// ChapterTable is the table for the chapter table above.
var ChapterTable = table.New(ChapterTableMeta)

// define struct Chapter for table coursez.chapter
type Chapter struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ModuleID    string `db:"moduleId"`
	CourseID    string `db:"courseId"`
	Sequence    int    `db:"sequence"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	IsActive    bool   `db:"is_active"`
}
