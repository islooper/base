# go-base

## 目录结构说明

### util

工具函数

### dao目录说明
```
dao
|-- db.go
|-- mongo.go
|-- redis.go
```

* db.go: gorm支持的数据库(MySQL, PostgreSQL, SQLite, SQL Server)通用dao以及注入函数
* [ ] mongo.go: mongodb通用dao以及注入函数(待定)
* [ ] redis.go: redis通用dao以及注入函数(待定)


### sharedmiddleware

api通用中间件模块， 建议一个中间件一个文件

### svc

资源池中的资源初始化函数

### sharedconfig

公用配置结构体
