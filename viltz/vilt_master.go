package viltz

import "github.com/scylladb/gocqlx/v2/table"

// create table viltz.vilt_master (
//	id varchar,
// 	lsp_id varchar,
// 	course_id varchar,
// 	no_of_learners bigint,
// 	trainers set<varchar>,
// 	moderators set<varchar>,
// 	course_start_date bigint,
// 	course_end_date bigint,
// 	curriculum text,
//  is_end_date_decided boolean,
//  is_moderator_decided boolean,
//  is_start_date_decided boolean,
//  is_trainer_decided boolean,
//  pricing_type varchar
//  price_per_seat bigint
//  currency varchar
//  tax_percent float
//  is_registration_open boolean
//  is_booking_open boolean
//  max_registrations int
//  registration_end_date bigint
//  booking_start_date bigint
//  booking_end_date bigint
//  registration_publish_by varchar
//  registration_publish_on bigint
//  booking_publish_on bigint
//  booking_publish_by varchar
//  registration_start_date bigint
// 	created_at bigint,
// 	updated_at bigint,
// 	created_by varchar,
// 	updated_by varchar,
//  status varchar,
// 	primary key((course_id), created_at)
// 	)
// 	with clustering order by (created_at DESC);

var ViltMasterMeta = table.Metadata{
	Name: "vilt_master",
	Columns: []string{
		"id",
		"lsp_id",
		"course_id",
		"no_of_learners",
		"trainers",
		"moderators",
		"course_start_date",
		"course_end_date",
		"curriculum",
		"is_end_date_decided",
		"is_moderator_decided",
		"is_start_date_decided",
		"is_trainer_decided",
		"pricing_type",
		"price_per_seat",
		"currency",
		"tax_percent",
		"is_registration_open",
		"is_booking_open",
		"max_registrations",
		"registration_end_date",
		"booking_start_date",
		"booking_end_date",
		"registration_publish_by",
		"registration_publish_on",
		"booking_publish_on",
		"booking_publish_by",
		"registration_start_date",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"course_id",
	},
	SortKey: []string{
		"created_at",
	},
}

var ViltMasterTable = table.New(ViltMasterMeta)

type ViltMaster struct {
	Id                    string   `db:"id"`
	LspId                 string   `db:"lsp_id"`
	CourseId              string   `db:"course_id"`
	NoOfLearners          int64    `db:"no_of_learners"`
	Trainers              []string `db:"trainers"`
	Moderators            []string `db:"moderators"`
	CourseStartDate       int64    `db:"course_start_date"`
	CourseEndDate         int64    `db:"course_end_date"`
	Curriculum            string   `db:"curriculum"`
	IsEndDateDecided      bool     `db:"is_end_date_decided"`
	IsModeratorDecided    bool     `db:"is_moderator_decided"`
	IsStartDateDecided    bool     `db:"is_start_date_decided"`
	IsTrainerDecided      bool     `db:"is_trainer_decided"`
	PriceType             string   `db:"pricing_type"`
	PricePerSeat          int64    `db:"price_per_seat"`
	Currency              string   `db:"currency"`
	TaxPercent            float32  `db:"tax_percent"`
	IsRegistrationOpen    bool     `db:"is_registration_open"`
	IsBookingOpen         bool     `db:"is_booking_open"`
	MaxRegistrations      int64    `db:"max_registrations"`
	RegistrationEndDate   int64    `db:"registration_end_date"`
	BookingStartDate      int64    `db:"booking_start_date"`
	BookingEndDate        int64    `db:"booking_end_date"`
	RegistrationPublishBy string   `db:"registration_publish_by"`
	RegistrationPublishOn int64    `db:"registration_publish_on"`
	BookingPublishOn      int64    `db:"booking_publish_on"`
	BookingPublishBy      string   `db:"booking_publish_by"`
	RegistrationStartDate int64    `db:"registration_start_date"`
	CreatedAt             int64    `db:"created_at"`
	CreatedBy             string   `db:"created_by"`
	UpdatedAt             int64    `db:"updated_at"`
	UpdatedBy             string   `db:"updated_by"`
	Status                string   `db:"status"`
}
