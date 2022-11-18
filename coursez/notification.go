package coursez

//define struct for sending notification to users
type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Token string `json:"token"`
}
