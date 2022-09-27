package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.cat_main (
//    id varchar,
//    name varchar,
//    description varchar,
//    image_bucket varchar,
//    image_url varchar,
//    code varchar,
//    is_active boolean,
//    created_at bigint,
//    updated_at bigint,
//    created_by varchar,
//    updated_by varchar,
//    lsps
//		set
//		< varchar >,
//    PRIMARY KEY (id)
//);
//
var CatMainMeta = table.Metadata{
	Name: "cat_main",
	Columns: []string{
		"id",
		"name",
		"description",
		"image_bucket",
		"image_url",
		"code",
		"is_active",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"lsps",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

var CatMainTable = table.New(CatMainMeta)

type CatMain struct {
	ID          string   `db:"id"`
	Name        string   `db:"name"`
	Description string   `db:"description"`
	ImageBucket string   `db:"image_bucket"`
	ImageURL    string   `db:"image_url"`
	Code        string   `db:"code"`
	IsActive    bool     `db:"is_active"`
	CreatedAt   int64    `db:"created_at"`
	UpdatedAt   int64    `db:"updated_at"`
	CreatedBy   string   `db:"created_by"`
	UpdatedBy   string   `db:"updated_by"`
	LspIDs      []string `db:"lsps"`
}
