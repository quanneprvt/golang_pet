package Models

type Game struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Played      float64 `json:"played"`
}
