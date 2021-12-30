package contracts

// Owner is the owner of a course
type Owner struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// CourseMeta is the metadata for a course
type CourseMeta struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Author         string   `json:"author"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	Category       string   `json:"category"`
	Tags           []string `json:"tags"`
	CourseID       string   `json:"course_id"`
	ExpertiseLevel string   `json:"expertise_level"`
	Languages      []string `json:"languages"`
	License        string   `json:"license"`
	Thumbnail      string   `json:"thumbnail"`
	PreviewLink    string   `json:"preview_link"`
	Owner          Owner    `json:"owner"`
}
