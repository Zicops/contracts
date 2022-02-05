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
		"summary",
		"instructor",
		"imageBucket",
		"image",
		"previewVideoBucket",
		"previewVideo",
		"tileImageBucket",
		"tileImage",
		"owner",
		"duration",
		"expertise_level",
		"language",
		"benefits",
		"created_at",
		"updated_at",
		"type",
		"prequisites",
		"goodFor",
		"mustFor",
		"related_skills",
		"publish_date",
		"expiry_date",
		"quality_control_check_reqd",
		"approvers",
		"created_by",
		"updated_by",
		"status",
		"is_active",
		"is_display",
		"expected_completion_time",
		"category",
		"sub_category",
		"sub_categories",

	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

// CourseTable is the table for the course table above.
var CourseTable = table.New(CourseTableMeta)

type SubCat struct {
	Name string `db:"name"`
	Rank int    `db:"rank"`
}

// define struct Course for table coursez.course
type Course struct {
	ID                 string   `db:"id"`
	Name               string   `db:"name"`
	Description        string   `db:"description"`
	Summary            string   `db:"summary"`
	Instructor         string   `db:"instructor"`
	ImageBucket        string   `db:"imageBucket"`
	Image              string   `db:"image"`
	PreviewVideoBucket string   `db:"previewVideoBucket"`
	PreviewVideo       string   `db:"previewVideo"`
	TileImageBucket    string   `db:"tileImageBucket"`
	TileImage          string   `db:"tileImage"`
	Owner              string   `db:"owner"`
	Duration           int      `db:"duration"`
	ExpertiseLevel     string   `db:"expertise_level"`
	Language           []string `db:"language"`
	Benefits           []string `db:"benefits"`
	CreatedAt          int64    `db:"created_at"`
	UpdatedAt          int64    `db:"updated_at"`
	Type               string   `db:"type"`
	Prequisites        []string `db:"prequisites"`
	GoodFor            []string `db:"goodFor"`
	MustFor            []string `db:"mustFor"`
	RelatedSkills      []string `db:"related_skills"`
	PublishDate        int64    `db:"publish_date"`
	ExpiryDate         int64    `db:"expiry_date"`
	QARequired         bool     `db:"quality_control_check_reqd"`
	Approvers          []string `db:"approvers"`
	CreatedBy          string   `db:"created_by"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
	IsActive           bool     `db:"is_active"`
	IsDisplay          bool     `db:"is_display"`
	ExpectedCompletion int64    `db:"expected_completion_time"`
	Category           string   `db:"category"`
	SubCategory        string   `db:"sub_category"`
	SubCategories      []SubCat `db:"sub_categories"`
}
