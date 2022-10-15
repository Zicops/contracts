package userz

import "github.com/scylladb/gocqlx/table"

//create table userz.users (
//	id varchar,
//	firstname varchar,
//	lastname varchar,
//	email varchar,
//	status varchar,
//	role varchar,
//	is_verified boolean,
//	gender varchar,
//	is_active boolean,
//	created_at bigint,
//	updated_at bigint,
//	created_by varchar,
//	updated_by varchar,
//	PRIMARY KEY (id)
//  );

var UserTableMeta = table.Metadata{
	Name: "users",
	Columns: []string{
		"id",
		"first_name",
		"last_name",
		"email",
		"status",
		"role",
		"is_verified",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_active",
		"gender",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{
		"created_at",
	},
}

var UserTable = table.New(UserTableMeta)

type User struct {
	ID          string `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Status      string `db:"status"`
	Role        string `db:"role"`
	IsVerified  bool   `db:"is_verified"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	Gender      string `db:"gender"`
	CreatedBy   string `db:"created_by"`
	UpdatedBy   string `db:"updated_by"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	PhotoBucket string `db:"photo_bucket"`
	PhotoURL    string `db:"photo_url"`
}
