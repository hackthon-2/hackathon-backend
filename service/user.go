package service

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hackthon/database/mysql"
	"hackthon/model"
	"hackthon/util"
	"os"
)

var UpdateProfileError = errors.New("update profile error")

func UploadAvatar(userId uint, fileName string) error {
	bucketName := os.Getenv("BUCKET_NAME")
	endPoint := os.Getenv("ENDPOINT")
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	config := util.OSSConfig{
		BucketName:      bucketName,
		EndPoint:        endPoint,
		AccessKeySecret: accessKeySecret,
		AccessKeyId:     accessKeyId,
	}
	bucket, err := config.GetOSSBucket()
	if err != nil {
		return err
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	err = bucket.PutObjectFromFile("avatars/"+fileName, "./avatars/"+fileName, storageType, objectAcl)
	if err != nil {
		return err
	}
	err = os.Remove("./avatars/" + fileName)
	if err != nil {
		return err
	}
	row, err := database.UpdateUserById(userId, &model.User{Avatar: "https://oss.onesnowwarrior.cn/avatars/" + fileName})
	if row != 1 {
		return UpdateProfileError
	}
	return err
}
