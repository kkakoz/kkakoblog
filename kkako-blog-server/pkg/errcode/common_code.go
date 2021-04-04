package errcode

var (
	Success = NewError(200, "ok")
)

var (
	InvalidParams = NewError(40001, "参数错误")
	NoAccountErr  = NewError(40002, "登陆失败,用户不存在")
	AccountExistErr = NewError(40003, "账户已经存在")
	PasswordErr = NewError(40004, "登录失败,密码错误")
	UserLoginErr = NewError(40005, "用户未登录,请登录后操作")
)

var (
	ServerErr            = NewError(50001, "服务器内部错误")
	NotFound             = NewError(50002, "未找到")
	AuthNotExist         = NewError(50003, "没有找到对应的appkey")
	AuthTokenErr         = NewError(50004, "token校验错误，请重新登录")
	AuthTokenTimeOut     = NewError(50005, "token校验超时")
	AuthTokenGenerateErr = NewError(50006, "token生成失败")
	ToomanyRequest       = NewError(50007, "请求过多")
	FindErr              = NewError(50008, "查找失败")
	FirstErr             = NewError(50009, "查找单个失败")
	CountErr             = NewError(50010, "获取数量失败")
	UpdateErr            = NewError(50011, "更新失败")
	AddErr               = NewError(50012, "添加失败")
	DeleteErr = NewError(50013, "删除失败")

)
