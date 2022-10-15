package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_notes (
//	id varchar,
//	user_id varchar,
//	user_lsp_id varchar,
//	course_id varchar,
//	module_id varchar,
//	details text,
//	status varchar,
//	is_active boolean,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//

var UserNotesTableMeta = table.Metadata{
	Name: "user_notes",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"course_id",
		"module_id",
		"topic_id",
		"details",
		"status",
		"sequence",
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
	SortKey: []string{
		"created_at",
	},
}

var UserNotesTable = table.New(UserNotesTableMeta)

type UserNotes struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	UserLspID string `db:"user_lsp_id"`
	CourseID  string `db:"course_id"`
	ModuleID  string `db:"module_id"`
	TopicID   string `db:"topic_id"`
	Details   string `db:"details"`
	Status    string `db:"status"`
	Sequence  int    `db:"sequence"`
	IsActive  bool   `db:"is_active" json:"is_active"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedBy string `db:"created_by"`
	UpdatedBy string `db:"updated_by"`
}
