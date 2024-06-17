package dao

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"tiktok/cmd/user/dal/db"
	"time"
)

type UserDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *db.User) error {
	err := dao.Model(&db.User{}).Create(&user).Error
	return err
}

func (dao *UserDao) FindUserByName(name string) (*db.User, error) {
	var retUser *db.User
	err := dao.DB.Model(&db.User{}).Where("username = ?", name).First(&retUser).Error //Db改dao
	return retUser, err
}

func (dao *UserDao) UpdateDate(user *db.User) error {
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := dao.DB.Save(user).Error //Db改dao
	return err
}

func (dao *UserDao) FindUserByUid(uid string) (*db.User, error) {
	var retUser *db.User
	u, _ := strconv.ParseInt(uid, 10, 64)
	err := dao.DB.Model(&db.User{}).Where("uid = ?", u).First(&retUser).Error //Db改dao
	return retUser, err
}

func (dao *UserDao) UploadAvatar(url string, user *db.User) (err error) {
	//user, _ := dao.FindUserByUid(strconv.FormatInt(uid, 10))
	user.AvatarUrl = url
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = dao.DB.Save(&user).Error
	return
}
