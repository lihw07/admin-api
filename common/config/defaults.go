package config

// RedisConfig Redis服务器配置
var RedisConfig = &redis{
	Network:  "tcp",
	Server:   "",
	Password: "",
	DB:       0,
}

// DatabaseConfig 数据库配置
var DbConfig = &db{
	Charset:    "utf8",
	Port:       3306,
	UnixSocket: false,
}

// SystemConfig 系统公用配置
var SystemConfig = &system{
	Debug:       false,
	Mode:        "master",
	Listen:      ":5212",
	ProxyHeader: "X-Forwarded-For",
}

// CORSConfig 跨域配置
var CORSConfig = &cors{
	AllowOrigins:     []string{"UNSET"},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
	AllowHeaders:     []string{"Cookie", "X-Cr-Policy", "Authorization", "Content-Length", "Content-Type", "X-Cr-Path", "X-Cr-FileName"},
	AllowCredentials: false,
	ExposeHeaders:    nil,
	SameSite:         "Default",
	Secure:           false,
}

var ImageSetting = &imageSettings{
	UploadDir: "/admin-api/upload/",
	ImageHost: "http://localhost:2000",
}

var LogConfig = &log{
	Path:  "./log",
	Name:  "sys",
	Model: "file",
}

var OptionOverwrite = map[string]interface{}{}
