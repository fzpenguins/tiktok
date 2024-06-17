package like

import (
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/rpc"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
)

func (s *LikeService) ListLikeVideo(req *interaction.ListLikeReq) (resp []*video.Video, err error) {
	var wg sync.WaitGroup
	uid, err := strconv.ParseInt(req.GetUid(), 10, 64)
	if err != nil {
		return nil, errors.WithMessage(err, errno.ParseFailed)
	}

	key := cache.GetVideoLikeFromUser(uid)
	var length []string
	pos := 0
	for {

		length = cache.RedisClient.ZRevRangeByScore(s.ctx, key, &redis.ZRangeBy{
			Min:    "1",
			Max:    "1",
			Offset: int64(pos),
			Count:  req.GetPageSize(),
		}).Val()
		if cache.RedisClient.ZCard(s.ctx, key).Val() <= req.GetPageSize() {
			break
		}
		pos += len(length)
		if int64(pos) >= req.GetPageNum()*req.GetPageSize() {
			break
		}
		log.Println("pos=", pos)
	}

	var vids []int64
	resp = make([]*video.Video, len(length))

	for _, v := range length {
		c, _ := strconv.ParseInt(v, 10, 64)
		vids = append(vids, c)
	}
	log.Println("vids=", vids)
	ret, err := rpc.GetVideoInfo(s.ctx, vids)
	if err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	for i, value := range ret.Items.Items {
		wg.Add(1)
		go func(value *video.Video, index int) {
			defer wg.Done()
			vinfo := &video.Video{
				Vid:         value.Vid, // strconv.FormatInt(value.Vid, 10),
				Uid:         value.Uid, //strconv.FormatInt(value.Uid, 10),
				VideoUrl:    value.VideoUrl,
				CoverUrl:    value.CoverUrl,
				Title:       value.Title,
				Description: value.Description,
				VisitCount:  value.VisitCount,
				CreatedAt:   value.CreatedAt,
				UpdatedAt:   value.UpdatedAt,
				DeletedAt:   value.DeletedAt,
			}
			log.Println("vinfo=", vinfo)
			t, err := s.GetVideoInfo(&interaction.GetVideoInfoRequest{Vid: strconv.FormatInt(value.Vid, 10)})
			if err != nil {
				return
			}
			vinfo.LikeCount, vinfo.CommentCount = t.LikeCount, t.CommentCount
			resp[index] = vinfo
		}(value, i)
	}
	wg.Wait()
	return
}
