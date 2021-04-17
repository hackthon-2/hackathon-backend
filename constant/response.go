package constant

type ResponseCode int

const (
	CODE_100 ResponseCode = 100 //参数错误
	CODE_101 ResponseCode = 101 //无效token
	CODE_102 ResponseCode = 102 //token过期
	CODE_103 ResponseCode = 103 //token刷新失败
	CODE_104 ResponseCode = 104 //token错误
	SUCCESS  ResponseCode = 200 //成功
	CODE_201 ResponseCode = 201 //服务器未获取到相应参数
	CODE_202 ResponseCode = 202 //用户名已被使用
	CODE_203 ResponseCode = 203 //密码错误
	CODE_204 ResponseCode = 204 //用户名不存在
	CODE_205 ResponseCode = 205 //密码为空
	CODE_206 ResponseCode = 206 //无权限访问
	CODE_207 ResponseCode = 207 //获取数据失败
	CODE_301 ResponseCode = 301 //创建日记失败
	CODE_302 ResponseCode = 302 //更新日记失败
	CODE_303 ResponseCode = 303 //删除日记失败
	CODE_304 ResponseCode = 304 //创建待办失败
	CODE_305 ResponseCode = 305 //更新待办失败
	CODE_306 ResponseCode = 306 //删除待办失败
	CODE_307 ResponseCode = 307 //获取统计数据失败
	CODE_308 ResponseCode = 308 //创建监督失败
	CODE_309 ResponseCode = 309 //更新监督失败
	CODE_310 ResponseCode = 310 //获取监督失败
	CODE_311 ResponseCode = 311 //删除监督失败
	CODE_401 ResponseCode = 401 //上传文件失败
	CODE_402 ResponseCode = 402 //更新用户个人资料失败
	CODE_403 ResponseCode = 403 //拉取用户个人资料失败
	CODE_404 ResponseCode = 404 //页面未找到
	CODE_500 ResponseCode = 500 //服务器错误
)

var codeTextMap = map[ResponseCode]string{
	CODE_100: "参数错误",
	CODE_101: "无效token",
	CODE_102: "token过期",
	CODE_103: "刷新token失败",
	CODE_104: "token错误",
	SUCCESS:  "成功",
	CODE_201: "参数错误",
	CODE_202: "用户名已被使用",
	CODE_203: "密码与用户名不匹配",
	CODE_204: "用户名不存在",
	CODE_205: "密码不能为空",
	CODE_206: "无权限访问",
	CODE_207: "获取数据失败",
	CODE_301: "创建日记失败",
	CODE_302: "更新日记失败",
	CODE_303: "删除日记失败",
	CODE_304: "创建待办失败",
	CODE_305: "更新待办失败",
	CODE_306: "删除待办失败",
	CODE_307: "获取统计数据失败",
	CODE_308: "创建监督失败",
	CODE_309: "更新监督失败",
	CODE_310: "获取监督失败",
	CODE_311: "删除监督失败",
	CODE_401: "上传文件失败",
	CODE_402: "更新用户个人资料失败",
	CODE_403: "获取用户个人资料失败",
	CODE_404: "页面不存在",
	CODE_500: "服务器内部错误",
}

func GetCodeText(code ResponseCode) string {
	if value, ok := codeTextMap[code]; ok {
		return value
	}
	return "Unknown code text"
}
