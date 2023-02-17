package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// create table vendorz.profile (
//     pf_id varchar,
//     vendor_id varchar,
//     type varchar,
//     first_name varchar,
//     last_name varchar,
//     email varchar,
//     phone_number varchar,
//     photo text,
//     photo_bucket text,
//     photo_url text,
//     description text,
//     languages set<varchar>,
//     sme_expertise set<varchar>,
//     classroom_expertise set<varchar>,
//     experience set<varchar>,
//     experience_years varchar,
//     is_speaker boolean,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((pf_id, vendor_id), type)
// );

var VendorProfileMeta = table.Metadata{
	Name: "profile",
	Columns: []string{
		"pf_id",
		"vendor_id",
		"type",
		"first_name",
		"last_name",
		"email",
		"phone_number",
		"photo",
		"photo_bucket",
		"photo_url",
		"description",
		"languages",
		"sme_expertise",
		"classroom_expertise",
		"experience",
		"experience_years",
		"is_speaker",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"status",
	},
	PartKey: []string{
		"pf_id",
		"vendor_id",
	},
	SortKey: []string{
		"type",
	},
}

var VendorProfileTable = table.New(VendorProfileMeta)

type VendorProfile struct {
	PfId               string   `db:"pf_id"`
	VendorId           string   `db:"vendor_id"`
	Type               string   `db:"type"`
	FirstName          string   `db:"first_name"`
	LastName           string   `db:"last_name"`
	Email              string   `db:"email"`
	Phone              string   `db:"phone_number"`
	Photo              string   `db:"photo"`
	PhotoBucket        string   `db:"photo_bucket"`
	PhotoURL           string   `db:"photo_url"`
	Description        string   `db:"description"`
	Languages          []string `db:"languages"`
	SMEExpertise       []string `db:"sme_expertise"`
	ClassroomExpertise []string `db:"classroom_expertise"`
	Experience         []string `db:"experience"`
	ExperienceYears    string   `db:"experience_years"`
	IsSpeaker          bool     `db:"is_speaker"`
	CreatedAt          int64    `db:"created_at"`
	CreatedBy          string   `db:"created_by"`
	UpdatedAt          int64    `db:"updated_at"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
}
