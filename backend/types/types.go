package types

type Participant struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type ActivityPreview struct {
	Title  string `json:"title"`
	Points int    `json:"points"`
}

type Activity struct {
	Title       string `json:"title"`
	Points      int    `json:"points"`
	Description string `json:"description"`
}

type Submission struct {
}
