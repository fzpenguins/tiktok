package db

type Video struct {
	Vid          int64  `json:"vid" gorm:"primaryKey;autoincrement"`
	Uid          int64  `json:"uid"`
	VideoUrl     string `json:"video_url"`
	CoverUrl     string `json:"cover_url"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	VisitCount   int64  `json:"visit_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
