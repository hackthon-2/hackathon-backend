package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	Username string `json:"username" gorm:"<-;type:VARCHAR(30) not NULL;collate:utf8mb4_unicode_ci;unique;index:idx_name,unique,type:btree,length:30"`
	Password string `json:"password" gorm:"<-;type:VARCHAR(100) not NULL"`
	Email    string `json:"email" gorm:"<-;type:VARCHAR(100) not NULL;collate:utf8mb4_unicode_ci;index:idx_email,type:btree,length:100"`
	Avatar   string `json:"avatar" gorm:"<-;type:VARCHAR(256) not NULL;collate:utf8mb4_unicode_ci;index:idx_avatar,type:btree,length:200;default:https://oss.onesnowwarrior.cn/avatars/default.jpg"`
}

// Diary 日记数据
type Diary struct {
	ID uint `json:"id" gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	//用户ID，通过tokenVerify后的claim中获得
	UserID   uint   `json:"user_id" gorm:"<-:create;type:INT UNSIGNED not NULL;index:idx_userId,type:btree"`
	Question string `json:"question" gorm:"<-;type:VARCHAR(100) not NULL;collate:utf8mb4_unicode_ci;index:idx_question,type:btree,length:100"`
	Text     string `json:"text" gorm:"<-;type:TEXT not NULL;collate:utf8mb4_unicode_ci;index:idx_text,type:btree,length:400"`
	Time     string `json:"time" gorm:"<-;type:DATE not NULL;collate:utf8mb4_unicode_ci;index:idx_time,type:btree,sort:desc"`
}

// Todo 待办和完成的列表
type Todo struct {
	ID     uint   `gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	UserID uint   `gorm:"<-:create;type:INT UNSIGNED not NULL;index:idx_userId,type:btree"`
	Header string `gorm:"<-;type:VARCHAR(100) not NULL;collate:utf8mb4_unicode_ci;index:idx_header,type:btree,length:100"`
	//这个存json到数据库里面，多个todoItem用/隔开，提取出来用split出[]string，然后遍历Unmarshal成[]ToDoItem
	TodoItem string `gorm:"<-;type:TEXT not NULL;collate:utf8mb4_unicode_ci;index:idx_todoItems,type:btree,length:450"`
	Time     string `gorm:"<-;type:DATE not NULL;collate:utf8mb4_unicode_ci;index:idx_time,type:btree,sort:desc"`
}

// Watch 监督的内容，待办
type Watch struct {
	ID           uint   `gorm:"primaryKey;<-:create;type:INT UNSIGNED not NULL auto_increment"`
	Content      string `gorm:"<-;type:TEXT;collate:utf8mb4_unicode_ci;index:idx_content,type:btree,length:400"`
	Username     string `gorm:"<-:create;collate:utf8mb4_unicode_ci;type:VARCHAR(30) not NULL;index:idx_name,type:btree,length:30"`
	UserID       uint   `gorm:"<-:create;type:INT UNSIGNED not NULL;index:idx_userId,type:btree"`
	Time         uint   `gorm:"<-:create;type:TINYINT UNSIGNED not NULL;index:idx_time,type:btree"`
	FinishedTime uint   `gorm:"<-;type:TINYINT UNSIGNED not NULL;index:idx_finishedTime,type:btree"`
	WatcherID    uint   `gorm:"<-:create;type:INT UNSIGNED not NULL;index:idx_watcherId,type:btree"`
	Watcher      string `gorm:"<-:create;collate:utf8mb4_unicode_ci;type:VARCHAR(30) not NULL;unique;index:idx_watcher,type:btree,length:30"`
}
type Profile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
