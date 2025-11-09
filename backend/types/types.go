package types

type Participant struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Activity struct {
	Title  string `json:"title"`
	Points int    `json:"points"`
	Link   string `json:"link"`
}
