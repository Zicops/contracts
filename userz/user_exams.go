package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_exam_attempts (
//	id varchar,
//	user_id varchar,
//	user_lsp_id varchar,
//	user_cp_id varchar,
//	exam_id varchar,
//	attempt_no bigint,
//	attempt_status varchar,
//	attempt_start_time bigint,
//	attempt_duration bigint,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//

var UserExamAttemptsTableMeta = table.Metadata{
	Name: "user_exam_attempts",
	Columns: []string{
		"id",
		"user_id",
		"user_lsp_id",
		"user_cp_id",
		"exam_id",
		"attempt_no",
		"attempt_status",
		"attempt_start_time",
		"attempt_duration",
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

var UserExamAttemptsTable = table.New(UserExamAttemptsTableMeta)

type UserExamAttempts struct {
	ID               string `db:"id"`
	UserID           string `db:"user_id"`
	UserLspID        string `db:"user_lsp_id"`
	UserCpID         string `db:"user_cp_id"`
	UserCmID         string `db:"user_cm_id"`
	ExamID           string `db:"exam_id"`
	AttemptNo        int64  `db:"attempt_no"`
	AttemptStatus    string `db:"attempt_status"`
	AttemptStartTime int64  `db:"attempt_start_time"`
	AttemptDuration  string `db:"attempt_duration"`
	CreatedAt        int64  `db:"created_at"`
	UpdatedAt        int64  `db:"updated_at"`
	CreatedBy        string `db:"created_by"`
	UpdatedBy        string `db:"updated_by"`
}
