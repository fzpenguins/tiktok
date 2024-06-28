package service

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"tiktok/cmd/picture/dal/db"
	"tiktok/cmd/picture/dal/db/dao"
	"tiktok/cmd/picture/dal/milvus"
	"tiktok/cmd/picture/rpc"
	"tiktok/kitex_gen/picture"
	"tiktok/proto"
)

func (s *PictureService) SearchByImage(req *picture.SearchByImageRequest) (images []*db.Image, err error) {
	var ids []int64
	resp, err := http.Get(req.Url)
	if err != nil {
		return nil, errors.WithMessage(err, "GetUrlByID failed")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "GetUrlByID failed")
	}

	vector, err := rpc.GetImageVector(s.ctx, &proto.ImageRequest{Image: data})
	if err != nil {
		return nil, errors.WithMessage(err, "GetTextVector failed")
	}
	ids, err = milvus.Search(s.ctx, vector)
	if err != nil {
		return nil, errors.WithMessage(err, "Search failed")
	}
	imageDao := dao.NewImageDao(s.ctx)
	images = make([]*db.Image, len(ids))
	for index, id := range ids {
		pic, err := imageDao.GetURLByPid(id)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
		temp := &db.Image{
			Pid: pic.Pid,
			Url: pic.Url,
		}
		images[index] = temp
	}
	return images, nil
}
