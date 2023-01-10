package coursez

import "github.com/scylladb/gocqlx/v2/table"

// create table coursez.discussion (
//     discussion_id varchar,
//     course_id varchar,
//     reply_id varchar,
//     user_id varchar,
//     time_stamp bigint,
//     content text,
//     module text,
//     chapter text,
//     topic text,
//     topic text,
//     likes set< varchar >,
//     dislikes set< varchar >,
//     is_anonymous boolean,
//     is_pinned boolean,
//     is_announcement boolean,
//     reply_count int,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((course_id, discussion_id), created_at)
// )

var DiscussionTableMeta = table.Metadata{
	Name: "discussion",
	Columns: []string{
		"discussion_id",
		"course_id",
		"reply_id",
		"user_id",
		"time_stamp",
		"content",
		"module",
		"chapter",
		"topic",
		"likes",
		"dislikes",
		"is_anonymous",
		"is_pinned",
		"is_announcement",
		"reply_count",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
		"status",
	},
	PartKey: []string{
		"discussion_id",
		"course_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var DiscussionTable = table.New(DiscussionTableMeta)

type Discussion struct {
	DiscussionId   string   `db:"discussion_id"`
	CourseId       string   `db:"course_id"`
	ReplyId        string   `db:"reply_id"`
	UserId         string   `db:"user_id"`
	Time           int64    `db:"time_stamp"`
	Content        string   `db:"content"`
	Module         string   `db:"module"`
	Chapter        string   `db:"chapter"`
	Topic          string   `db:"topic"`
	Likes          []string `db:"likes"`
	Dislike        []string `db:"dislikes"`
	IsAnonymous    bool     `db:"is_anonymous"`
	IsPinned       bool     `db:"is_pinned"`
	IsAnnouncement bool     `db:"is_announcement"`
	ReplyCount     int      `db:"reply_count"`
	CreatedBy      string   `db:"created_by"`
	CreatedAt      int64    `db:"created_at"`
	UpdatedBy      string   `db:"updated_by"`
	UpdatedAt      int64    `db:"updated_at"`
	Status         string   `db:"status"`
}
