package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/interaction/dal/db"
)

type LikeDao struct {
	*gorm.DB
}

func NewLikeDao(ctx context.Context) *LikeDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &LikeDao{NewDBClient(ctx)}
}

func (dao *LikeDao) CreateLike(req *db.Like) error {
	err := dao.DB.Model(&db.Like{}).Create(&req).Error
	return err
}

func (dao *LikeDao) GetVideoLikeCount(vid int64) (int64, error) {
	var count int64
	err := dao.DB.Model(&db.Like{}).Where("vid = ?", vid).Count(&count).Error
	return count, err
}

func (dao *LikeDao) GetUserLikeVideoExist(uid, vid int64) bool {
	var count int64
	dao.DB.Model(&db.Like{}).Where("vid = ? AND uid = ?", vid, uid).Count(&count)
	return count == 0
}

func (dao *LikeDao) GetUserLikeCommentExist(uid, cid int64) bool {
	var count int64
	dao.DB.Model(&db.Like{}).Where("cid = ? AND uid = ?", cid, uid).Count(&count)
	return count == 0
}

func (dao *LikeDao) GetCommentLikeCount(cid int64) (int64, error) {
	var count int64
	err := dao.DB.Model(&db.Like{}).Where("cid = ?", cid).Count(&count).Error
	return count, err
}

func (dao *LikeDao) DeleteVideoLike(uid, vid int64) error {
	return dao.DB.Model(&db.Like{}).Delete("uid = ? AND vid = ?", uid, vid).Error
}

func (dao *LikeDao) DeleteCommentLike(uid, cid int64) error {
	return dao.DB.Model(&db.Like{}).Delete("uid = ? AND cid = ?", uid, cid).Error
}
