package cache

import (
	"context"
	"strconv"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
	"time"

	"github.com/pkg/errors"
)

func VisitCount(ctx context.Context, vid int64) int64 {
	countStr, _ := RedisClient.Get(ctx, GetVideoVisitCountKey(vid)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

func AddVisitCount(ctx context.Context, video *db.Video) (err error) {
	txn := RedisClient.TxPipeline()

	err = txn.Incr(ctx, GetVideoVisitCountKey(video.Vid)).Err()
	if err != nil {
		return err
	}
	err = txn.ZIncrBy(ctx, "visit_count", 1, strconv.Itoa(int(video.Vid))).Err()
	if err != nil {
		return err
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func IsExistsVideoInfo(ctx context.Context, vid string) bool {
	return RedisClient.Exists(ctx, VideoInfoKey(vid)).Val() == 1
}

func AddVideoInfo(ctx context.Context, vid string, data []byte) (err error) {
	err = RedisClient.Set(ctx, VideoInfoKey(vid), string(data), 5*time.Minute).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddVideoInfo failed")
	}
	return
}

func GetVideoInfo(ctx context.Context, vid string) (data string, err error) {
	data, err = RedisClient.Get(ctx, VideoInfoKey(vid)).Result()
	if err != nil {
		return "", errors.Wrap(err, "cache.GetVideoInfo failed")
	}
	return
}

func GetVideoRank(ctx context.Context, videoRequest *video.PopularReq) (vids []int64, offset int64, cnt int64, err error) {

	offset = videoRequest.GetPageNum() * videoRequest.GetPageSize()
	cnt, _ = RedisClient.ZCard(ctx, "visit_count").Result()

	Stringcmds, err := RedisClient.ZRevRange(ctx, "visit_count", 0, -1).Result()
	if err != nil {
		return nil, 0, 0, err
	}

	if offset >= int64(len(Stringcmds)) {
		return []int64{}, 0, 0, err
	}

	for i := offset; i-offset < videoRequest.GetPageSize() && i < cnt; i++ {
		vid, _ := strconv.ParseInt(Stringcmds[i], 10, 64)
		vids = append(vids, vid)
	}

	return
}

func GetUid(ctx context.Context, req *video.SearchReq) (string, error) {
	return RedisClient.Get(ctx, req.GetUsername()).Result()
}
