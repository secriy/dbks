# 数据库课程设计
> 企业网站管理系统
> 项目结构仿照[ACGFATE](https://github.com/secriy/acgfate)
> 未使用任何ORM

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