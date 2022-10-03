package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_exam_progress (
//	id varchar,
//	user_id varchar,
//	user_ea_id varchar,
//	user_lsp_id varchar,
//	user_cp_id varchar,
//	sr_no bigint,
//	question_id varchar,
//	question_type varchar,
//	answer text,
//	q_attempt_status varchar,
//	total_time_spent bigint,
//	correct_answer text,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//

var UserExamProgressTableMeta = table.Metadata{
	Name: "user_exam_progress",
	Columns: []string{
		"id",
		"user_id",
		"user_ea_id",
		"user_lsp_id",
		"user_cp_id",
		"sr_no",
		"question_id",
		"question_type",
		"answer",
		"q_attempt_status",
		"total_time_spent",
		"correct_answer",
		"section_id",
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

var UserExamProgressTable = table.New(UserExamProgressTableMeta)

type UserExamProgress struct {
	ID             string `db:"id"`
	UserID         string `db:"user_id"`
	UserEaID       string `db:"user_ea_id"`
	UserLspID      string `db:"user_lsp_id"`
	UserCpID       string `db:"user_cp_id"`
	SrNo           int64  `db:"sr_no"`
	QuestionID     string `db:"question_id"`
	QuestionType   string `db:"question_type"`
	Answer         string `db:"answer"`
	QAttemptStatus string `db:"q_attempt_status"`
	TotalTimeSpent int64  `db:"total_time_spent"`
	CorrectAnswer  string `db:"correct_answer"`
	SectionID      string `db:"section_id"`
	CreatedAt      int64  `db:"created_at"`
	UpdatedAt      int64  `db:"updated_at"`
	CreatedBy      string `db:"created_by"`
	UpdatedBy      string `db:"updated_by"`
}
