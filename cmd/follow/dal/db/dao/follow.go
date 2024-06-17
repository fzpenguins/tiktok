package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/follow/dal/db"
	"tiktok/pkg/errno"
)

type FollowDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func NewFollowDao(ctx context.Context) *FollowDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &FollowDao{NewDBClient(ctx)}
}

func (dao *FollowDao) CreateFollow(f *db.Follow) error {
	return dao.DB.Model(&db.Follow{}).Create(&f).Error
}

func (dao *FollowDao) DeleteFollow(fromUid, toUid string) error {
	return dao.DB.Model(&db.Follow{}).Delete("from_uid = ? AND to_uid = ?", fromUid, toUid).Error
}

func (dao *FollowDao) GetFollowing(fromUid int64) ([]*db.Follow, error) {
	var list []*db.Follow
	err := dao.DB.Model(&db.Follow{}).Where("from_uid = ?", fromUid).Find(&list).Error
	return list, err
}

func (dao *FollowDao) GetFollower(toUid int64) ([]*db.Follow, error) {
	var list []*db.Follow
	err := dao.DB.Model(&db.Follow{}).Where("to_uid = ?", toUid).Find(&list).Error
	return list, err
}

//func (dao *FollowDao) GetFriends(fromUid, toUid string) {
//	dao.DB.Model(&db.Follow{}).Where("to_uid = ? AND from_uid = ?", toUid, fromUid)
//}

func (dao *FollowDao) IsFollowerExist(fromUid, toUid int64) error {
	var cnt int64
	err := dao.DB.Model(&db.Follow{}).Where("from_uid = ? AND to_uid = ?", fromUid, toUid).Count(&cnt).Error
	if cnt != 0 {
		return errno.FollowExistedError
	}
	return err
}

func (dao *FollowDao) IsFollowed(fromUid, toUid int64) bool {
	var cnt int64
	dao.DB.Model(&db.Follow{}).Where("from_uid = ? AND to_uid = ?", fromUid, toUid).Count(&cnt)
	return cnt != 0
}