package config

import (
	"github.com/mlogclub/simple/sqls"
)

var Instance *Config

type Config struct {
	Env        string // 环境：prod、dev
	BaseUrl    string // base url
	Port       string // 端口
	LogFile    string // 日志文件
	StaticPath string // 静态文件目录
	IpDataPath string // IP数据文件

	// 数据库配置
	DB sqls.DbConfig

	// 阿里云oss配置
	Uploader struct {
		Enable    string
		AliyunOss struct {
			Host          string
			Bucket        string
			Endpoint      string
			AccessId      string
			AccessSecret  string
			StyleSplitter string
			StyleAvatar   string
			StylePreview  string
			StyleSmall    string
			StyleDetail   string
		}
		Local struct {
			Host string
			Path string
		}
	}

	// 百度SEO相关配置
	// 文档：https://ziyuan.baidu.com/college/courseinfo?id=267&page=2#h2_article_title14
	BaiduSEO struct {
		Site  string
		Token string
	}

	// 神马搜索SEO相关
	// 文档：https://zhanzhang.sm.cn/open/mip
	SmSEO struct {
		Site     string
		UserName string
		Token    string
	}

	// smtp
	Smtp struct {
		Host     string
		Port     string
		Username string
		Password string
		SSL      bool
	}

	// es
	Es struct {
		Url   string
		Index string
	}
}