package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.user_course_register (
// 	id varchar,
// 	course_id varchar,
// 	user_id varchar,
// 	registration_date bigint,
// 	invoice varchar,
// 	status varchar,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updated_by varchar,
// 	primary key((course_id), updated_at)
// 	)
// 	with clustering order by (updated_at desc);

var UserCourseRegisterMeta = table.Metadata{
	Name: "user_course_register",
	Columns: []string{
		"id",
		"course_id",
		"user_id",
		"registration_date",
		"invoice",
		"status",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	},
	PartKey: []string{
		"course_id",
	},
	SortKey: []string{
		"updated_at",
	},
}

var UserCourseRegisterTable = table.New(UserCourseRegisterMeta)

type UserCourseRegister struct {
	Id               string `db:"id"`
	CourseId         string `db:"course_id"`
	UserId           string `db:"user_id"`
	RegistrationDate int64  `db:"registration_date"`
	Invoice          string `db:"invoice"`
	Status           string `db:"status"`
	CreatedAt        int64  `db:"created_at"`
	CreatedBy        string `db:"created_by"`
	UpdatedAt        int64  `db:"updated_at"`
	UpdatedBy        string `db:"updated_by"`
}
