package model

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var ScretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

func UpLoadFile(file multipart.File, fileSize int64)  (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	//putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(AccessKey, ScretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS: false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil{
		return err.Error(), errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}