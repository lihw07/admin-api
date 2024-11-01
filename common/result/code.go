package result

// 状态码状态语工具类
// @author lihw07

// Codes 定义的状态
type Codes struct {
	SUCCESS                    uint
	FAILED                     uint
	Message                    map[uint]string
	NOAUTH                     uint
	AUTHFORMATERROR            uint
	INVALIDTOKEN               uint
	MissingLoginParameter      uint
	VerificationCodeHasExpired uint
	CAPTCHANOTTRUE             uint
	PASSWORDNOTTRUE            uint
	STATUSISENABLE             uint
	ROLENAMEALREADYEXISTS      uint
	MENUISEXIST                uint
	DELSYSMENUFAILED           uint
	DEPTISEXIST                uint
	DEPTISDISTRIBUTE           uint
	POSTALREADYEXISTS          uint
	FILEUPLOADERROR            uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:                    200,
	FAILED:                     501,
	NOAUTH:                     403,
	AUTHFORMATERROR:            405,
	INVALIDTOKEN:               406,
	MissingLoginParameter:      407,
	VerificationCodeHasExpired: 408,
	CAPTCHANOTTRUE:             409,
	PASSWORDNOTTRUE:            410,
	STATUSISENABLE:             411,
	ROLENAMEALREADYEXISTS:      412,
	MENUISEXIST:                413,
	DELSYSMENUFAILED:           414,
	DEPTISEXIST:                415,
	DEPTISDISTRIBUTE:           416,
	POSTALREADYEXISTS:          417,
	FILEUPLOADERROR:            427,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                    "成功",
		ApiCode.FAILED:                     "失败",
		ApiCode.NOAUTH:                     "请求头中auth为空",
		ApiCode.AUTHFORMATERROR:            "请求头中auth格式有误",
		ApiCode.INVALIDTOKEN:               "无效的Token或者登录过期,请重新登录！",
		ApiCode.MissingLoginParameter:      "缺少登录参数",
		ApiCode.VerificationCodeHasExpired: "验证码已失效",
		ApiCode.CAPTCHANOTTRUE:             "验证码不正确，请重新输入",
		ApiCode.PASSWORDNOTTRUE:            "密码不正确",
		ApiCode.STATUSISENABLE:             "您的账号已被停用,请联系管理员",
		ApiCode.ROLENAMEALREADYEXISTS:      "角色名称或权限字符已存在，请重新输入",
		ApiCode.MENUISEXIST:                "菜单名称已存在，请重新输入",
		ApiCode.DELSYSMENUFAILED:           "菜单已分配，不能删除",
		ApiCode.DEPTISEXIST:                "部门名称已存在，请重新输入",
		ApiCode.DEPTISDISTRIBUTE:           "部门已分配，不能删除",
		ApiCode.POSTALREADYEXISTS:          "岗位名称或岗位编码已存在，请重新输入",
		ApiCode.FILEUPLOADERROR:            "文件上传错误",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
