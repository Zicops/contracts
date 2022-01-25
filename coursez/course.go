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
//        created_at bigint,
//        updated_at bigint,
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
		"takeAways",
		"created_at",
		"updated_at",
		"type",
		"prequisites",
		"goodFor",
		"mustFor",
		"created_by",
		"updated_by",
		"status",
		"is_deleted",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// CourseTable is the table for the course table above.
var CourseTable = table.New(CourseTableMeta)

// define struct Course for table coursez.course
type Course struct {
	ID                 string   `db:"id"`
	Name               string   `db:"name"`
	Description        string   `db:"description"`
	Instructor         string   `db:"instructor"`
	ImageBucket        string   `db:"imageBucket"`
	Image              string   `db:"image"`
	PreviewVideoBucket string   `db:"previewVideoBucket"`
	PreviewVideo       string   `db:"previewVideo"`
	Owner              string   `db:"owner"`
	Duration           int      `db:"duration"`
	Level              string   `db:"level"`
	Language           []string `db:"language"`
	TakeAways          []string `db:"takeAways"`
	CreatedAt          int64    `db:"created_at"`
	UpdatedAt          int64    `db:"updated_at"`
	Type               string   `db:"type"`
	Prequisites        []string `db:"prequisites"`
	GoodFor            []string `db:"goodFor"`
	MustFor            []string `db:"mustFor"`
	CreatedBy          string   `db:"created_by"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
	IsDeleted          bool     `db:"is_deleted"`
}
