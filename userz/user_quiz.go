package userz

import "github.com/scylladb/gocqlx/table"

// create table userz.user_quiz_attempts (
// 	id varchar,
// 	user_id varchar,
// 	user_cp_id varchar,
// 	user_cm_id varchar,
// 	topic_id varchar,
// 	quiz_id varchar,
// 	quiz_attempt bigint,
// 	result varchar,
// 	is_active boolean,
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
// 	PRIMARY KEY (id)
//   );

var UserQuizAttemptsTableMeta = table.Metadata{
	Name: "user_quiz_attempts",
	Columns: []string{
		"id",
		"user_id",
		"user_cp_id",
		"user_cm_id",
		"topic_id",
		"quiz_id",
		"quiz_attempt",
		"result",
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

var UserQuizAttemptsTable = table.New(UserQuizAttemptsTableMeta)

type UserQuizAttempts struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	UserCpID    string `db:"user_cp_id"`
	UserCmID    string `db:"user_cm_id"`
	TopicID     string `db:"topic_id"`
	QuizID      string `db:"quiz_id"`
	QuizAttempt int64  `db:"quiz_attempt"`
	Result      string `db:"result"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
}
