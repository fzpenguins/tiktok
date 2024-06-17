package db

type Follow struct {
	Fid       int64  `json:"fid" gorm:"primaryKey"`
	FromUid   int64  `json:"from_uid"`
	ToUid     int64  `json:"to_uid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
