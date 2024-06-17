package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"strconv"
	"sync"
	"tiktok/cmd/follow/dal/cache"
	"tiktok/cmd/follow/dal/db/dao"
	"tiktok/cmd/follow/rpc"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"
)

func (s *FollowService) FriendList(req *follow.ListFriendReq) (*follow.ListFriendResp, error) {
	uid, _ := strconv.ParseInt(req.GetUid(), 10, 64)
	resp := new(follow.ListFriendResp)
	resp.Data = &follow.UserInfoData{}

	offset := req.GetPageNum() * req.GetPageSize()

	log.Println("90909")

	rets, err := cache.FriendListAction(s.ctx, req.Uid)
	if err != nil {
		return resp, errors.WithMessage(err, errno.GetInfoError)
	} else if len(rets) == 0 { //数据库中访问
		list, err := dao.NewFollowDao(s.ctx).GetFollowing(uid)
		if err != nil {
			return resp, err
		}
		log.Println(list)
		for _, d := range list {
			err := dao.NewFollowDao(s.ctx).IsFollowerExist(d.ToUid, uid)
			if errors.Is(err, gorm.ErrRecordNotFound) { //找不到
				continue
			} else if err != nil {
				return nil, err
			}
			err = cache.AddFollow(s.ctx, strconv.FormatInt(d.ToUid, 10), req.GetUid())
			if err != nil {
				return nil, err
			}
			rets = append(rets, strconv.FormatInt(d.ToUid, 10))
		}
		log.Println("two,", rets)
	}
	log.Println("ret = ", rets)
	if int(offset) >= len(rets) {
		resp.Data = &follow.UserInfoData{
			Items: nil,
			Total: 0,
		}
		return resp, errno.ParamError
	}
	resp.Data = &follow.UserInfoData{

		Total: int64(len(rets)),
	}
	var wg sync.WaitGroup
	for i, ret := range rets {
		if i >= int(offset) && i < int(req.GetPageSize()+offset) {
			wg.Add(1)
			go func(value string, index int) {
				defer wg.Done()
				uInfo, err := rpc.GetUserInfo(s.ctx, value)
				if err != nil {
					return
				}
				resp.Data.Items = append(resp.Data.Items, uInfo)
			}(ret, i)
		}
	}
	wg.Wait()

	return resp, nil
}
