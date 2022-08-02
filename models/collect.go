package models

// Collect represents a collect
type Collect struct {
	CollectID string `json:"collect_id"`
	UserID    string `json:"user_id"`
	URL       string `json:"url"`
}
