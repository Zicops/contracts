package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.topic_classroom (
// 	id varchar,
// 	topic_id varchar,
// 	trainer set<varchar>,
// 	moderator set<varchar>,
// 	training_start_time bigint,
// 	training_end_time bigint,
// 	duration bigint,
// 	breaktime varchar,
// 	language varchar,
// 	is_screen_share_enabled boolean,
// 	is_chat_enabled boolean,
// 	is_microphone_enabled boolean,
// 	is_qa_enabled boolean,
// 	is_camera_enabled boolean,
// 	is_override_config boolean,
// 	created_at bigint,
// 	created_by varchar,
// 	updated_at bigint,
// 	updated_by varchar,
// 	status varchar,
// 	PRIMARY KEY((id, topic_id))
// 	);

var TopicClassroomMeta = table.Metadata{
	Name: "topic_classroom",
	Columns: []string{
		"id",
		"topic_id",
		"trainer",
		"moderator",
		"training_start_time",
		"training_end_time",
		"duration",
		"breaktime",
		"language",
		"is_screen_share_enabled",
		"is_chat_enabled",
		"is_microphone_enabled",
		"is_qa_enabled",
		"is_camera_enabled",
		"is_override_config",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"id",
		"topic_id",
	},
}

var TopicClassroomTable = table.New(TopicClassroomMeta)

type TopicClassroom struct {
	Id                   string   `db:"id"`
	TopicId              string   `db:"topic_id"`
	Trainer              []string `db:"trainer"`
	Moderator            []string `db:"moderator"`
	TrainingStartTime    int64    `db:"training_start_time"`
	TrainingEndTime      int64    `db:"training_end_time"`
	Duration             int64    `db:"duration"`
	Breaktime            string   `db:"breaktime"`
	Language             string   `db:"language"`
	IsScreenShareEnabled bool     `db:"is_screen_share_enabled"`
	IsChatEnabled        bool     `db:"is_chat_enabled"`
	IsMicrophoneEnabled  bool     `db:"is_microphone_enabled"`
	IsQAEnabled          bool     `db:"is_qa_enabled"`
	IsCameraEnabled      bool     `db:"is_camera_enabled"`
	IsOverrideConfig     bool     `db:"is_override_config"`
	CreatedAt            int64    `db:"created_at"`
	CreatedBy            string   `db:"created_by"`
	UpdatedAt            int64    `db:"updated_at"`
	UpdatedBy            string   `db:"updated_by"`
	Status               string   `db:"status"`
}
