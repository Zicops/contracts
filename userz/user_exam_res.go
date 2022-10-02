package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.user_exam_results (
//	id varchar,
//	user_id varchar,
//	user_ea_id varchar,
//	user_score bigint,
//	correct_answers bigint,
//	wrong_answers bigint,
//	result_status varchar,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );
//

var UserExamResultsTableMeta = table.Metadata{
	Name: "user_exam_results",
	Columns: []string{
		"id",
		"user_id",
		"user_ea_id",
		"user_score",
		"correct_answers",
		"wrong_answers",
		"result_status",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	},
	PartKey: []string{
		"user_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserExamResultsTable = table.New(UserExamResultsTableMeta)

type UserExamResults struct {
	ID             string `db:"id"`
	UserID         string `db:"user_id"`
	UserEaID       string `db:"user_ea_id"`
	UserScore      int64  `db:"user_score"`
	CorrectAnswers int64  `db:"correct_answers"`
	WrongAnswers   int64  `db:"wrong_answers"`
	ResultStatus   string `db:"result_status"`
	CreatedAt      int64  `db:"created_at"`
	UpdatedAt      int64  `db:"updated_at"`
	CreatedBy      string `db:"created_by"`
	UpdatedBy      string `db:"updated_by"`
}
