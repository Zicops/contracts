package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_bookmarks (
//	id varchar,
//	user_id varchar,
//	user_lsp_id varchar,
//	course_id varchar,
//	module_id varchar,
//	user_cp_id varchar,
//	name varchar,
//	time_stamp bigint,
//	is_active boolean,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//

var UserBookmarksTableMeta = table.Metadata{
	Name: "user_bookmarks",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"course_id",
		"module_id",
		"topic_id",
		"user_cp_id",
		"name",
		"time_stamp",
		"is_active",
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

var UserBookmarksTable = table.New(UserBookmarksTableMeta)

type UserBookmarks struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	UserLspID string `db:"user_lsp_id"`
	CourseID  string `db:"course_id"`
	ModuleID  string `db:"module_id"`
	TopicID   string `db:"topic_id"`
	UserCPID  string `db:"user_cp_id"`
	Name      string `db:"name"`
	TimeStamp string `db:"time_stamp"`
	IsActive  bool   `db:"is_active" json:"is_active"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
}
