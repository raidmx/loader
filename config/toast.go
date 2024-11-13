package config

// Toast is a new UI based popup that appears in the top right corner of the screen of the player
// by default.
type Toast struct {
	// Title of the toast
	Title string `json:"title"`

	// Description of the toast
	Content string `json:"content"`
}
