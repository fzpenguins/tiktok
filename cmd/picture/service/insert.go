package service

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"tiktok/cmd/picture/dal/db"
	"tiktok/cmd/picture/dal/db/dao"
	"tiktok/cmd/picture/dal/milvus"
	"tiktok/cmd/picture/dal/minio"
	"tiktok/cmd/picture/rpc"
	"tiktok/kitex_gen/picture"
	"tiktok/proto"
)

func (s *PictureService) Insert(req *picture.InsertRequest) (image *db.Image, err error) {
	//现将数据上传到minio
	finalStr, err := minio.UploadImage(req.Url)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert upload data error")
	}

	//保存到mysql
	//确保上传成功后再写url保存到mysql，保持数据一致

	imageDao := dao.NewImageDao(s.ctx)
	image = &db.Image{
		Url: finalStr,
	}
	image, err = imageDao.CreateImage(image)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert create image error")
	}

	resp, err := http.Get(req.Url)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert create image error")
	}
	defer resp.Body.Close()

	imageByte, err := ioutil.ReadAll(resp.Body)

	//rpc调用获取向量数据
	vector, err := rpc.GetImageVector(s.ctx, &proto.ImageRequest{Image: imageByte})
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert get vector error")
	}
	err = milvus.InsertVector(s.ctx, vector, image.Pid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert insert vector error")
	}
	return image, nil
}
