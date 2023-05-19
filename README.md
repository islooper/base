# go-base 🤦🏿

```go
 go version >= 1.19
```

## 目录结构说明

### util

工具函数

### dao目录说明
```
dao
|-- db.go
|-- redis.go
```

* db.go: gorm支持的数据库(MySQL, PostgreSQL, SQLite, SQL Server)通用dao以及注入函数
* [ ] redis.go: redis通用dao以及注入函数(待定)


### httplib

http工具包（但是好像有bug

### locklib

基于etcd的分布式锁，不好用。

### mq 
不好用

### sharedmiddleware

api通用中间件模块， 建议一个中间件一个文件

### svc

资源池中的资源初始化函数

### sharedconfig

公用配置结构体

### util
各类集合