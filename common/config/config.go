package config

import (
	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
)

// system 系统通用配置
type system struct {
	Mode          string `validate:"eq=master|eq=slave"`
	Listen        string `validate:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
	GracePeriod   int    `validate:"gte=0"`
	ProxyHeader   string `validate:"required_with=Listen"`
}

// database 数据库
type db struct {
	Username    string
	Password    string
	Host        string
	Database    string
	TablePrefix string
	Port        int
	Charset     string
	UnixSocket  bool
	MaxIdle     int
	MaxOpen     int
}

// redis 配置
type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	DB       int
}

// imageSettings图片上传配置
type imageSettings struct {
	UploadDir string
	ImageHost string
}

// log日志
type log struct {
	Path  string
	Name  string
	Model string
}

// 跨域配置
type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	SameSite         string
	Secure           bool
}

// 声明变量
var cfg *ini.File

// 1.初始化配置文件
func Init(path string) {
	var err error

	// 2.加载配置文件
	cfg, err = ini.Load(path)
	if err != nil {
		panic(err)
	}

	// 3.加载默认配置
	sections := map[string]interface{}{
		"Db":           DbConfig,
		"System":       SystemConfig,
		"Redis":        RedisConfig,
		"Log":          LogConfig,
		"ImageSetting": ImageSetting,
		"CORS":         CORSConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			panic(err)
		}
	}

	// 4.映射数据库配置覆盖
	for _, key := range cfg.Section("OptionOverwrite").Keys() {
		OptionOverwrite[key.Name()] = key.Value()
	}

}

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := cfg.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}

	return nil
}
