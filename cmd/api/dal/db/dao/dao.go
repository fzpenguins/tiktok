package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/api/dal/db"
	"tiktok/pkg/constants"
	"time"
)

type MsgDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func GetMsgDao(ctx context.Context) *MsgDao {
	return &MsgDao{NewDBClient(ctx)}
}

//分页的详情由server端设定即可，无需client传入

func (dao *MsgDao) GetHistoryFromSingleChat(pageNum int, from, to string) (msgs []*db.Message, err error) {
	err = dao.DB.Model(&db.Message{}).Where("from_uid = ? AND to_uid = ? AND type = ?", from, to, 0).Limit(constants.PageSize).
		Offset(pageNum * constants.PageSize).Find(&msgs).Error
	return
}

func (dao *MsgDao) GetHistoryFromGroupChat(pageNum int, to string) (msgs []*db.Message, err error) {
	err = dao.DB.Model(&db.Message{}).Where("to_uid = ? AND type = ?", to, 1).Limit(constants.PageSize).
		Offset(pageNum * constants.PageSize).Find(&msgs).Error
	return
}

func (dao *MsgDao) GetUnreadFromSingleChat(pageNum int, from, to string) (msgs []*db.Message, err error) {
	err = dao.DB.Model(&db.Message{}).Where("from_uid = ? AND to_uid = ? AND read_tag = ?", from, to, false).Limit(constants.PageSize).
		Offset(pageNum * constants.PageSize).Find(&msgs).Error
	return
}

func (dao *MsgDao) StoreSingleChatMsg(from, to, content string, readTag bool) error {
	msg := &db.Message{
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		DeletedAt: "",
		FromUid:   from,
		ToUid:     to,
		Type:      constants.SingleChat,
		Content:   content,
		ReadTag:   readTag,
	}
	return dao.DB.Model(&db.Message{}).Create(&msg).Error
}

func (dao *MsgDao) TurnToRead(from, to string) error {
	return dao.Model(&db.Message{}).Where("from_uid = ? AND to_uid = ?", from, to).Update("read_tag", true).Error
}

func (dao *MsgDao) CreateMsg(m *db.Message) error {
	return dao.Model(&db.Message{}).Create(&m).Error
}
