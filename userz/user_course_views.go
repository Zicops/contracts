package userz

import "github.com/scylladb/gocqlx/v2/table"

// create table userz.user_course_views (
// 	course_id varchar,
// 	date_value varchar,
// 	time bigint,
// 	created_at bigint,
// 	users set<varchar>,
// 	primary key((course_id), created_at)
// 	)
// 	with clustering order by (created_at desc);

var UserCourseViewsMeta = table.Metadata{
	Name: "user_course_views",
	Columns: []string{
		"course_id",
		"date_value",
		"time",
		"created_at",
		"users",
	},
	PartKey: []string{
		"course_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserCourseViewsTable = table.New(UserCourseViewsMeta)

type UserCourseViews struct {
	CourseId  string   `db:"course_id"`
	DateValue string   `db:"date_value"`
	Time      int64    `db:"time"`
	CreatedAt int64    `db:"created_at"`
	Users     []string `db:"users"`
}
