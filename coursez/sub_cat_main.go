package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.sub_cat_main (
//    id varchar,
//    name varchar,
//    description varchar,
//    parent_id varchar,
//    image_bucket varchar,
//    image_url varchar,
//    code varchar,
//    is_active boolean,
//    created_at bigint,
//    PRIMARY KEY (id)
//    updated_at bigint,
//    created_by varchar,
//    updated_by varchar,
//);
//

var SubCatMainMeta = table.Metadata{
	Name: "sub_cat_main",
	Columns: []string{
		"id",
		"name",
		"description",
		"parent_id",
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
		"is_active",
	},
	SortKey: []string{},
}

var SubCatMainTable = table.New(SubCatMainMeta)

type SubCatMain struct {
	ID          string   `db:"id"`
	Name        string   `db:"name"`
	Description string   `db:"description"`
	ParentID    string   `db:"parent_id"`
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
