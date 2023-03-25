package vendorz

import "github.com/scylladb/gocqlx/v2/table"

// create table vendorz.profile (
//     pf_id varchar,
//     vendor_id varchar,
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
//     content_development set<varchar>,
//     experience set<varchar>,
//     experience_years varchar,
//     is_speaker boolean,
//     lsp_id varchar,
//     sme boolean,
//     crt boolean,
//     cd boolean,
//     name set<varchar>,
//     created_at bigint,
//     created_by varchar,
//     updated_at bigint,
//     updated_by varchar,
//     status varchar,
//     PRIMARY KEY((pf_id, vendor_id))
// );

var VendorProfileMeta = table.Metadata{
	Name: "profile",
	Columns: []string{
		"pf_id",
		"vendor_id",
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
		"content_development",
		"experience",
		"experience_years",
		"is_speaker",
		"lsp_id",
		"sme",
		"crt",
		"cd",
		"name",
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
}

var VendorProfileTable = table.New(VendorProfileMeta)

type VendorProfile struct {
	PfId               string   `db:"pf_id"`
	VendorId           string   `db:"vendor_id"`
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
	ContentDevelopment []string `db:"content_development"`
	Experience         []string `db:"experience"`
	ExperienceYears    string   `db:"experience_years"`
	IsSpeaker          bool     `db:"is_speaker"`
	LspId              string   `db:"lsp_id"`
	Sme                bool     `db:"sme"`
	Crt                bool     `db:"crt"`
	Cd                 bool     `db:"cd"`
	Name               []string `db:"words"`
	CreatedAt          int64    `db:"created_at"`
	CreatedBy          string   `db:"created_by"`
	UpdatedAt          int64    `db:"updated_at"`
	UpdatedBy          string   `db:"updated_by"`
	Status             string   `db:"status"`
}
