package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_course_map (
// 	id varchar,
// 	user_id varchar,
// 	user_lsp_id varchar,
// 	course_id varchar,
// 	course_type varchar,
// 	added_by varchar,
// 	is_mandatory boolean,
// 	course_status varchar,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserCourseMapTableMeta = table.Metadata{
	Name: "user_course_map",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"course_id",
		"course_type",
		"added_by",
		"is_mandatory",
		"end_date",
		"course_status",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"user_id",
		"id",
	},
	SortKey: []string{},
}

var UserCourseTable = table.New(UserCourseMapTableMeta)

type UserCourse struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	UserLspID    string `db:"user_lsp_id"`
	CourseID     string `db:"course_id"`
	CourseType   string `db:"course_type"`
	AddedBy      string `db:"added_by"`
	IsMandatory  bool   `db:"is_mandatory"`
	EndDate      int64  `db:"end_date"`
	CourseStatus string `db:"course_status"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
	CreatedBy    string `db:"created_by"`
	UpdatedBy    string `db:"updated_by"`
}
