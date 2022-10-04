package coursez

import (
	"reflect"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
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
		"lsp_id",
		"publisher",
		"description",
		"summary",
		"instructor",
		"imagebucket",
		"image",
		"previewvideobucket",
		"previewvideo",
		"tileimagebucket",
		"tileimage",
		"owner",
		"duration",
		"expertise_level",
		"language",
		"benefits",
		"outcomes",
		"created_at",
		"updated_at",
		"type",
		"prequisites",
		"goodfor",
		"mustfor",
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
		"lsp_id",
		"id",
		"is_active",
	},
	SortKey: []string{},
}

// CourseTable is the table for the course table above.
var CourseTable = table.New(CourseTableMeta)

type SubCat struct {
	Name string `cql:"name"`
	Rank int    `cql:"rank"`
}

// define struct Course for table coursez.course
type Course struct {
	ID                 string   `db:"id"`
	Name               string   `db:"name"`
	LspId              string   `db:"lsp_id" json:"lsp_id"`
	Publisher          string   `db:"publisher"`
	Description        string   `db:"description"`
	Summary            string   `db:"summary"`
	Instructor         string   `db:"instructor"`
	ImageBucket        string   `db:"imagebucket"`
	Image              string   `db:"image"`
	PreviewVideoBucket string   `db:"previewvideobucket"`
	PreviewVideo       string   `db:"previewvideo"`
	TileImageBucket    string   `db:"tileimagebucket"`
	TileImage          string   `db:"tileimage"`
	Owner              string   `db:"owner"`
	Duration           int      `db:"duration"`
	ExpertiseLevel     string   `db:"expertise_level"`
	Language           []string `db:"language"`
	Benefits           []string `db:"benefits"`
	Outcomes           []string `db:"outcomes"`
	CreatedAt          int64    `db:"created_at"`
	UpdatedAt          int64    `db:"updated_at"`
	Type               string   `db:"type"`
	Prequisites        []string `db:"prequisites"`
	GoodFor            []string `db:"goodfor"`
	MustFor            []string `db:"mustfor"`
	RelatedSkills      []string `db:"related_skills"`
	PublishDate        string   `db:"publish_date"`
	ExpiryDate         string   `db:"expiry_date"`
	QARequired         bool     `db:"quality_control_check_reqd"`
	Approvers          []string `db:"approvers"`
	CreatedBy          string   `db:"created_by"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
	IsActive           bool     `db:"is_active" json:"is_active"`
	IsDisplay          bool     `db:"is_display"`
	ExpectedCompletion string   `db:"expected_completion_time"`
	Category           string   `db:"category"`
	SubCategory        string   `db:"sub_category"`
	SubCategories      []string `db:"sub_categories"`
}

// MarshalUDT implements UDTMarshaler.
func (u SubCat) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	f := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(u), name)
	return gocql.Marshal(info, f.Interface())
}

// UnmarshalUDT implements UDTUnmarshaler.
func (u *SubCat) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	f := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(u), name)
	return gocql.Unmarshal(info, data, f.Addr().Interface())
}
