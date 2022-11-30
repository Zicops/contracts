package notificationz

type Message struct {
	Notification Skeleton `json:"notification"`
	To           string   `json:"to"`
	CreatedAt    int64    `json:"created_at"`
}

type Skeleton struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Results struct {
	MessageId string `json:"message_ids"`
}

type RespBody struct {
	Multicast_id  int `json:"multicast_id"`
	Success       int `json:"success"`
	Failure       int `json:"failure"`
	Canonical_ids int `json:"canonical_id"`
	Results       []Results
}
