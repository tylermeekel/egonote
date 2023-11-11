package types

type Note struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Sharelink string `json:"sharelink,omitempty"`
}

func ValidateNote(n Note) bool {
	return true
}
