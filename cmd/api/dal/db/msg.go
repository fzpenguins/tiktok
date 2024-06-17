package db

type Message struct {
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	FromUid   string `json:"from_uid"`
	ToUid     string `json:"to_uid"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	ReadTag   bool   `json:"read_tag"`
}
