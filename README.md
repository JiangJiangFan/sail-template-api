# 初始化项目

```sh

go mod init sail-chat
go mod tidy

```

# 结构

```tree

代码结构
  bootstrap # 初始化文件，包括数据库、日志等
  config  # yaml 对应文件
  global # 全局配置
  middleware # 中间件
  models # 模型
  res # 自定义请求格式
  routers # 路由
  services # 服务层，逻辑处理
  storage # 日志文件
  types # 前端参数类型
  utils # 工具库
  sail-chat.go # 入口文件
```

# 运行

```bash
# 运行
go run sail-chat.go server
# 停止
go run sail-chat.go stop
```

# SQL

`create_at` 添加 CURRENT_TIMESTAMP 以使用自动添加值
`update_at` 添加 CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 以使用自动添加值
