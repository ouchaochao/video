# video
A project for looking job.

Finish api!

```shell
api
├── dbops # 操作数据库
│   ├── api.go # 增删改查用户、视频、评论
│   ├── api_test.go # 测试api
│   ├── conn.go # 连接数据库
│   └── internal.go # 增删查session
├── defs # 定义
│   ├── apidef.go # 结构体定义
│   └── errs.go # 错误定义
├── handlers
│   ├── auth.go # 验证session
│   ├── handlers.go # 读取请求并处理请求
│   └── response.go # 返回处理后的信息
├── main.go # 主函数入口
├── session 
│   └── ops.go # 产生session、验证session失效、删除session
└── utils
    └── uuid.go # 产生通用唯一识别码

```