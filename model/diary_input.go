package model

type DiaryInput struct {
	//问题内容
	Question string `form:"question"`
	//日记内容
	Text string `form:"text"`
	//提交日期
	Time string `form:"time"`
}

type UpdateDiaryInput struct {
	ID uint `form:"diary_id"`
	//问题内容
	Question string `form:"question"`
	//日记内容
	Text string `form:"text"`
	//提交日期
	Time string `form:"time"`
}
