package dao

import (
	"context"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"

	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) CreateVideo(v *db.Video) (*db.Video, error) {
	err := dao.DB.Model(&db.Video{}).Create(&v).Error
	return v, err
}

func (dao *VideoDao) FindVideosByTimeStr(timeStr string) (videoList []*db.Video, err error) {
	err = dao.DB.Model(&db.Video{}).Where("created_at >= ?", timeStr).Limit(50).Order("created_at desc").Find(&videoList).Error
	return
}

func (dao *VideoDao) GetPublishList(req *video.ListReq) (list []*db.Video, cnt int64, err error) {

	err = dao.DB.Model(&db.Video{}).Where("uid IN (?)", req.Uid).
		Order("created_at desc").Limit(int(req.PageSize)).
		Offset(int(req.PageNum * req.PageSize)).Find(&list).Count(&cnt).Error
	if err != nil {
		return nil, 0, err
	}
	if len(list) != 0 {
		return
	}
	return nil, 0, nil
}

func (dao *VideoDao) FindVideoByVid(vid []int64) (video []*db.Video, err error) {

	err = dao.DB.Model(&db.Video{}).Where("vid IN (?)", vid).Find(&video).Error
	return
}

//func (dao *VideoDao) GetVideosInCondition(ctx context.Context, req *video.SearchReq) {
//	query := dao.DB.Model(&db.Video{}).Where("description LIKE ? OR title LIKE ?", "%"+req.Keywords+"%", "%"+req.Keywords+"%")
//	if req.GetUsername() != "" {
//		uid, err := cache.GetUid(ctx, req)
//		dao.DB.Model(&db.Video{}).Where("uid = ?", uid)
//	}
//
//	if req.GetFromDate() > 0 {
//		FromDateString := time.Unix(req.GetFromDate(), 0).Format("2006-01-02 15:04:05")
//		query = query.Where("created_at >= ?", FromDateString)
//	}
//	if req.GetToDate() > 0 {
//		ToDateString := time.Unix(req.GetToDate(), 0).Format("2006-01-02 15:04:05")
//		query = query.Where("created_at >= ?", ToDateString)
//	}
//}

func (dao *VideoDao) GetVideoByKeyword(keyword string) *gorm.DB {
	query := dao.DB.Model(&db.Video{}).Where("description LIKE ? OR title LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	return query
}

func (dao *VideoDao) GetVideoByFromDate(fromTime string) *gorm.DB {
	return dao.DB.Model(&db.Video{}).Where("created_at >= ?", fromTime)
}

func (dao *VideoDao) GetVideoByToDate(toDate string) *gorm.DB {
	return dao.DB.Model(&db.Video{}).Where("created_at >= ?", toDate)
}

func (dao *VideoDao) GetVideoByUid(uid string) *gorm.DB {
	return dao.DB.Model(&db.Video{}).Where("uid = ?", uid)
}

func (dao *VideoDao) GetVideos(pageNum, pageSize int64) (list []*db.Video, count int64, err error) {
	list = make([]*db.Video, pageSize)
	err = dao.DB.Model(&db.Video{}).Limit(int(pageSize)).Offset(int(pageNum * pageSize)).
		Find(&list).Count(&count).Error
	return
}

func (dao *VideoDao) GetVideoByVid(vid int64) (*db.Video, error) {
	video := new(db.Video)
	err := dao.DB.Model(&db.Video{}).Where("vid = ?", vid).Find(&video).Error
	return video, err
}
