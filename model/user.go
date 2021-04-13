package model

type User struct {
	ID       uint   `gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	Username string `gorm:"<-;type:VARCHAR(30) not NULL;unique;index:idx_name,unique,type:btree,length:30"`
	Password string `gorm:"<-;type:VARCHAR(100) not NULL"`
	Email    string `gorm:"<-;type:VARCHAR(100) not NULL;index:idx_email,type:btree,length:100"`
	Avatar   string `gorm:"<-;type:VARCHAR(256) not NULL;index:idx_avatar,type:btree,length:200;default:https://tieba-simulating.oss-cn-hangzhou.aliyuncs.com/avatar/default.jpg"`
	Sex      uint   `gorm:"<-;type:TINYINT UNSIGNED"`
}
