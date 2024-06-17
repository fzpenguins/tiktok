package db

type Like struct {
	Lid       int64  `json:"lid" gorm:"primaryKey;autoincrement"`
	Vid       int64  `json:"vid"`
	Uid       int64  `json:"uid"`
	Cid       int64  `json:"cid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"delete_at"`
}
