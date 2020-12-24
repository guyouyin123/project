# project
# 介绍

随便写了点东西，后端gin+gorm,前端vue搭建了一个前后台分离的web网站。数据资源是毒app爬虫抓下来的，爬虫没有放在上面。

Gin文件夹是后端项目

Jeff_vue是前端项目

# 前台项目启动

在前端项目的根目录下：

cnpm install 重新构建项目依赖环境

没有依赖自己cnpm安装一下，然后启动完成啦

# 后台go项目启动

GO111MODULE="on"  打开go mod模式

GOPROXY=https://goproxy.io。设置代理

```
//数据库换上自己的
db, _ := gorm.Open("mysql", "root:admin123@/dududu?charset=utf8&parseTime=True&loc=Local")
```

go run main.go   编译并执行

