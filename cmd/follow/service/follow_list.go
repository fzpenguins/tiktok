package service

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
	"sync"
	"tiktok/cmd/follow/dal/cache"
	"tiktok/cmd/follow/dal/db/dao"
	"tiktok/cmd/follow/rpc"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"
)

func (s *FollowService) FollowList(req *follow.ListFollowingReq) (*follow.ListFollowingResp, error) {
	var list []string
	var err error
	resp := new(follow.ListFollowingResp)
	resp.Data = &follow.UserInfoData{}
	offset := req.GetPageNum() * req.GetPageSize()
	log.Println(111111)

	if cache.IsFollowingSetExist(s.ctx, req.Uid) {
		list, err = cache.GetFollowingList(s.ctx, req.GetPageNum(), req.GetPageSize(), req.GetUid())
		if err != nil {
			return resp, errors.WithMessage(err, errno.GetInfoError)
		} else if len(list) == 0 { //mysql查找
			uid, _ := strconv.ParseInt(req.Uid, 10, 64)
			likes, err := dao.NewFollowDao(s.ctx).GetFollowing(uid)
			if err != nil {
				return nil, errors.WithMessage(err, errno.QueryFailed)
			}
			for _, item := range likes {
				u := strconv.FormatInt(item.ToUid, 10)
				list = append(list, u)
				err = cache.AddFollow(s.ctx, strconv.FormatInt(item.FromUid, 10), strconv.FormatInt(item.ToUid, 10))
				if err != nil {
					continue
				}
			}
			log.Println("one")
		}
	} else {
		uid, _ := strconv.ParseInt(req.Uid, 10, 64)
		likes, err := dao.NewFollowDao(s.ctx).GetFollowing(uid)
		if err != nil {
			return nil, errors.WithMessage(err, errno.QueryFailed)
		}
		for _, item := range likes {
			u := strconv.FormatInt(item.ToUid, 10)
			list = append(list, u)
			err = cache.AddFollow(s.ctx, strconv.FormatInt(item.FromUid, 10), strconv.FormatInt(item.ToUid, 10))
			if err != nil {
				continue
			}
		}
		log.Println("two")
	}
	log.Println(list)
	log.Println(222)

	log.Println(1)

	if int(offset) >= len(list) {
		resp.Data = &follow.UserInfoData{
			Items: nil,
			Total: 0,
		}
		return resp, errno.ParamError
	}

	resp.Data = &follow.UserInfoData{

		Total: int64(len(list)),
	}
	var wg sync.WaitGroup
	for i, l := range list {
		if i >= int(offset) && i < int(req.GetPageSize()+offset) {
			wg.Add(1)
			go func(value string, index int) {
				defer wg.Done()
				ret, err := rpc.GetUserInfo(s.ctx, value)
				if err != nil {
					return
				}
				resp.Data.Items = append(resp.Data.Items, ret)
			}(l, i)
		}

		log.Println("i=", l)
	}
	wg.Wait()
	log.Println(resp)
	return resp, nil
}
