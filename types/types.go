package types

type Car struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	Make      string `json:"make"`
	Model     string `json:"model"`
	Year      string `json:"year"`
	ImageURL  string `json:"imageURL"`
	CreatedAt string `json:"createdAt"`
}
