package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.cohort_main (
//	id varchar,
//	lsp_id varchar,
//	code varchar,
//	name varchar,
//	type varchar,
//	description text,
//	is_active boolean,
//	status varchar,
//	imageBucket varchar,
//	imageUrl varchar,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//
var CohortMain = table.Metadata{
	Name: "cohort_main",
	Columns: []string{
		"id",
		"lsp_id",
		"code",
		"name",
		"type",
		"description",
		"is_active",
		"status",
		"imageBucket",
		"imageUrl",
		"size",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{},
}

var CohortTable = table.New(CohortMain)

type Cohort struct {
	ID          string `db:"id"`
	LspID       string `db:"lsp_id"`
	Code        string `db:"code"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	Description string `db:"description"`
	IsActive    bool   `db:"is_active"`
	Status      string `db:"status"`
	ImageBucket string `db:"imageBucket"`
	ImageUrl    string `db:"imageUrl"`
	Size        int    `db:"size"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
}