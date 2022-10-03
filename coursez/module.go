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
		"ischapter",
		"name",
		"description",
		"courseid",
		"lsp_id",
		"owner",
		"duration",
		"created_at",
		"updated_at",
		"level",
		"sequence",
		"setglobal",
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

// ModuleTable is the table for the module table above.
var ModuleTable = table.New(ModuleTableMeta)

// define struct Module for table coursez.module
type Module struct {
	ID          string `db:"id"`
	IsChapter   bool   `db:"ischapter"`
	Name        string `db:"name"`
	Description string `db:"description"`
	CourseID    string `db:"courseid"`
	LspId       string `db:"lsp_id" json:"lsp_id"`
	Owner       string `db:"owner"`
	Duration    int    `db:"duration"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	Level       string `db:"level"`
	Sequence    int    `db:"sequence"`
	SetGlobal   bool   `db:"setglobal"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}
