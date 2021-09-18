package Models

type User struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Password     string  `json:"password"`
	RecentPlayed float64 `json:"recent_played"`
}
