package dao

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/kitex_gen/interaction"
)

type CommentDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}

func (dao *CommentDao) CreateComment(comment *db.Comment) (err error) {
	err = dao.DB.Model(&db.Comment{}).Create(&comment).Error
	return
}

func (dao *CommentDao) GetCommentsByVid(req *interaction.ListCommentReq) ([]*db.Comment, error) {
	var list []*db.Comment
	err := dao.DB.Model(&db.Comment{}).Where("vid = ?", req.Vid).Limit(int(req.PageSize)).
		Offset(int(req.PageSize * req.PageNum)).Find(&list).Error
	log.Println("list = ", list)
	return list, err
}

func (dao *CommentDao) GetCommentsByCid(req *interaction.ListCommentReq) ([]*db.Comment, error) {
	var list []*db.Comment
	err := dao.DB.Model(&db.Comment{}).Where("parent_id = ?", req.Cid).Limit(int(req.PageSize)).
		Offset(int(req.PageSize * req.PageNum)).Find(&list).Error
	return list, err
}

func (dao *CommentDao) DeleteComment(req *interaction.DeleteReq) (*db.Comment, error) {
	var ret *db.Comment
	err := dao.DB.Model(&db.Comment{}).
		Where("vid = ? AND uid = ? AND cid = ?", req.GetVid(), req.GetUid(), req.GetCid()).
		First(&ret).Delete(&db.Comment{}).Error
	return ret, err
}

func (dao *CommentDao) GetCommentByCid(cid int64) (*db.Comment, error) {
	resp := new(db.Comment)
	var cnt int64
	err := dao.DB.Model(&db.Comment{}).Where("cid = ?", cid).Find(&resp).Count(&cnt).Error
	if cnt == 0 {
		return nil, errors.New("comment not found")
	}
	return resp, err
}
