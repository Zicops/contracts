package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.course_views (
////    lsp_id varchar,
////    date_value varchar,
////    users set <varchar>,
////    created_at bigint,
////    hours bigint,
////    PRIMARY KEY ((lsp_id, date_value), created_at)
////)
////
//WITH CLUSTERING ORDER BY (created_at DESC);

var CVMain = table.Metadata{
	Name: "course_views",
	Columns: []string{
		"lsp_id",
		"date_value",
		"users",
		"created_at",
		"hours",
	},
	PartKey: []string{
		"lsp_id",
		"date_value",
	},
	SortKey: []string{
		"created_at",
	},
}

var CVTable = table.New(CVMain)

type CourseView struct {
	LspId     string   `db:"lsp_id"`
	DateValue string   `db:"date_value"`
	Users     []string `db:"users"`
	CreatedAt int64    `db:"created_at"`
	Hours     int64    `db:"hours"`
}
