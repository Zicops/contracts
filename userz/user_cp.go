package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_course_progress (
// 	id varchar,
// 	user_id varchar,
// 	user_cm_id varchar,
// 	topic_id varchar,
// 	topic_type varchar,
// 	status varchar,
// 	time_stamp bigint,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserCourseProgressTableMeta = table.Metadata{
	Name: "user_course_progress",
	Columns: []string{
		"id",
		"user_id",
		"user_cm_id",
		"topic_id",
		"topic_type",
		"status",
		"video_progress",
		"time_stamp",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"user_id",
		"id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserCourseProgressTable = table.New(UserCourseProgressTableMeta)

type UserCourseProgress struct {
	ID            string `db:"id"`
	UserID        string `db:"user_id"`
	UserCmID      string `db:"user_cm_id"`
	TopicID       string `db:"topic_id"`
	TopicType     string `db:"topic_type"`
	Status        string `db:"status"`
	VideoProgress string `db:"video_progress"`
	TimeStamp     int64  `db:"time_stamp"`
	CreatedAt     int64  `db:"created_at"`
	UpdatedAt     int64  `db:"updated_at"`
	CreatedBy     string `db:"created_by"`
	UpdatedBy     string `db:"updated_by"`
}
