package types

type Participant struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Activity struct {
	Title       string `json:"title"`
	Points      int    `json:"points"`
	Description string `json:"description"`
}

// Status: "unreviewed" | "rejected" | "accepted"
//
// Image: []byte
type Submission struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Image  []byte `json:"image"`
}
