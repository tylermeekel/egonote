package types

type Note struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Sharelink string `json:"sharelink,omitempty"`
	UserID  int `json:"user_id,omitempty"`
}

func ValidateNote(n Note) bool {
	return true
}
