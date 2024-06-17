package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/user/dal/cache"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	minioclient "tiktok/cmd/user/dal/minio"
	"tiktok/config"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
	"tiktok/pkg/utils"
)

func (s *UserService) GetMfa(req *user.GetMFAReq) (*db.User, string, error) {
	usr, err := dao.NewUserDao(s.ctx).FindUserByUid(strconv.FormatInt(req.Uid, 10))
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.QueryFailed)
	}

	img, key, err := utils.GenerateQRCode(usr.Username, req.Uid)
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.GenerateFailed)
	}

	imgByBase64, err := utils.ImageToBase64(img)
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.ParseFailed)
	}

	decodedImg, err := base64.StdEncoding.DecodeString(imgByBase64)
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.ParseFailed)
	}

	buf := bytes.NewBuffer(decodedImg)
	storeQRCodePath := "qrcode/" + "uid:" + strconv.FormatInt(req.Uid, 10) + "user:" + usr.Username + ".png"
	_, err = minioclient.MinioClient.PutObject(context.Background(), config.BucketName, storeQRCodePath, buf, -1, minio.PutObjectOptions{
		ContentType: "image/png", // 设置内容类型为图片
	})
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.SaveFileFailed)
	}
	usr.Secret = key.Secret()
	usr.CodeUrl = storeQRCodePath
	err = cache.SetUserMFASecret(s.ctx, usr.Username, key.Secret())
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.SaveFileFailed)
	}
	err = cache.SetUserMFACodeUrl(s.ctx, usr.Username, storeQRCodePath)
	if err != nil {
		return nil, "", errors.WithMessage(err, errno.SaveFileFailed)
	}

	return usr, imgByBase64, nil
}
