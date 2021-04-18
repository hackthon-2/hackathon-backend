package service

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hackthon/database/mysql"
	"hackthon/model"
	"hackthon/util"
	"os"
)

var (
	UpdateProfileError = errors.New("update profile error")
	ListProfileError   = errors.New("list profile error")
)

// UploadAvatar 上传oss的函数
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
	//上传到oss成功，就把本地存的给删了
	err = os.Remove("./avatars/" + fileName)
	if err != nil {
		return err
	}
	//更新一下用户的头像地址
	row, err := database.UpdateUserById(userId, &model.User{Avatar: "https://oss.onesnowwarrior.cn/avatars/" + fileName})
	if row != 1 {
		return UpdateProfileError
	}
	return err
}
func ListProfile(userId uint) (model.Profile, error) {
	user, row, err := database.FindUserById(userId)
	if err != nil {
		return model.Profile{}, err
	}
	if row != 1 {
		return model.Profile{}, ListProfileError
	}
	return user, nil
}
func UpdateProfile(userId uint, input *model.UpdateUserInput) error {
	_, rows, _ := database.FindUserByUsername(input.Username)
	if rows != 0 {
		return ExistedUsername
	}
	var user model.User
	util.StructAssign(&user, input)
	row, err := database.UpdateUserById(userId, &user)
	if err != nil {
		return err
	}
	if row != 1 {
		return UpdateProfileError
	}
	return nil
}
