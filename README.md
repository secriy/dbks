# 数据库课程设计—企业网站管理系统

> 使用Golang的gin框架实现
>
> 手写SQL语句，未使用ORM
>
> @Author：[Secriy](https://github.com/secriy), [Devil](https://gitee.com/YJDevil), [cocop](https://gitee.com/cocop)

## 框架

1. [Gin](https://github.com/gin-gonic/gin): Go Web框架
2. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
3. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具
4. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件

## 项目结构

- api：封装了所有的接口，通过gin框架实现表单验证
- auth：权限控制文件
- conf：读取配置，设置环境变量
- database：连接数据库，初始化数据库表
- middleware：中间件
- model：数据库模型
- router：分发路由
- serializer：序列化器，将模型转换为json对象
- service：处理服务，实现增删查改
- util：小工具
- .env：配置文件
- main.go：程序入口

## 模型

- user：用户
- news：企业新闻
- offers：招聘信息
- products：产品信息
- info：企业信息

## 接口

- 公开接口
	- POST   /api/user/login — 用户登录
	- GET    /api/news/all — 查询所有新闻
	- GET    /api/offers/all — 查询所有招聘信息
	- GET    /api/products/all — 查询所有产品信息
	- GET    /api/info/all — 查询企业信息
- 普通用户权限接口
	- GET    /api/user/me — 查询用户个人资料
	- DELETE /api/user/logout — 用户登出
	- PUT    /api/user/pass — 修改密码
- 管理员权限接口
	- GET    /api/user/all — 查询所有用户
	- POST   /api/user/add — 增加用户
	- DELETE /api/user/del/:id — 删除用户
	- PUT    /api/user/update/:id — 更新用户
	- POST   /api/news/add — 增加新闻
	- DELETE /api/news/del/:id — 删除新闻
	- PUT    /api/news/update/:id — 更新新闻
	- POST   /api/offers/add — 增加招聘信息
	- DELETE /api/offers/del/:id — 删除招聘信息
	- PUT    /api/offers/update/:id — 更新招聘信息
	- POST   /api/products/add — 增加产品信息
	- DELETE /api/products/del/:id — 删除产品信息
	- PUT    /api/products/update/:id — 更新产品信息
	- DELETE /api/info/del/:option — 删除企业信息
	- PUT    /api/info/update — 更新产品信息

## 编译方法

1. 安装[Golang](https://golang.org/dl/)

2. 配置环境变量
	
	```
	GOPATH=D:\Environment\GoPath	// 并非程序安装文件夹，请任意指定其他文件夹
	GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on
	// 在path中配置
	D:\Environment\Go\bin	// 程序安装位置
	```

3. 执行`go build`，进行编译

## 运行

1. 创建数据库dbks（utf-8）
2. 在命令行下cd到项目目录，执行`./server.exe`