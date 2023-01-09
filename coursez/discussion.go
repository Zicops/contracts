package coursez

import "github.com/scylladb/gocqlx/v2/table"

// create table coursez.discussion (
// 	discussion_id varchar,
// 	course_id text,
// 	reply_id text,
// 	user text,
// 	time bigint,
// 	content text,
// 	current_topic topic_object,
// 	likes set< varchar >,
// 	dislikes set< varchar >,
// 	is_anonymous boolean,
// 	is_pinned boolean,
// 	is_announcement boolean,
// 	reply_count int,
// 	created_by text,
// 	created_at text,
// 	updated_by text,
// 	updated_at text,
// 	status text,
// 	PRIMARY KEY (id)
// );

var discussionTableMeta = table.Metadata{
	Name: "discussion",
	Columns: []string{
		"discussion_id",
		"course_id",
		"reply_id",
		"user",
		"time",
		"content",
		"current_topic",
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
		"discussionId",
		"courseId",
	},
	SortKey: []string{},
}

var discussionTable = table.New(discussionTableMeta)

type discussion struct {
	DiscussionId   string      `db:"discussionId"`
	CourseId       string      `db:""`
	ReplyId        string      `db:"replyId"`
	User           userdata    `db:"user_id"`
	Time           int         `db:"time"`
	Content        string      `db:"content"`
	CurrentTopic   topicObject `db:"current_topic"`
	Likes          []int       `db:"likes"`
	Dislike        []int       `db:"dislikes"`
	IsAnonymous    bool        `db:"is_anonymous"`
	IsPinned       bool        `db:"is_pinned"`
	IsAnnouncement bool        `db:"is_announcement"`
	ReplyCount     int         `db:"reply_count"`
	CreatedBy      string      `db:"created_by"`
	Created_at     string      `db:"created_at"`
	Updated_by     string      `db:"updated_by"`
	Updated_at     string      `db:"updated_at"`
	Status         string      `db:"status"`
}

type userdata struct {
	UserId     int
	FirstName  string
	LastName   string
	ProfileImg string
}

type topicObject struct {
	Module  string
	Chapter string
	Topic   string
	Time    string
}
