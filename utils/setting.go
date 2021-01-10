package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

//包初始化接口
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQinu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js67235")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("cyt1997511")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQinu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("ZinVMsu3OhP8Nl8UFuQqkBLaRTXFjzhV5ZIp6Ers")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("wXJH378QHaCIC5LPhVBzHHnmZ31VRRf68AFO2ciq")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("cytginbog")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("http://qmq2kmtqa.hd-bkt.clouddn.com/")
}
