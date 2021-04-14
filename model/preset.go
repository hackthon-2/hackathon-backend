package model

// Question 预设问题
type Question struct {
	ID   uint   `gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	Text string `gorm:"<-:create;type:TEXT not NULL;collate:utf8mb4_unicode_ci;index:idx_text,type:btree,length:400"`
}
