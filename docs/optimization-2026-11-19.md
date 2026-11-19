# 2026-11-19 扫描：问题与优化步骤

只读扫描本仓库。本文不记录密钥明文。总览见 [r2day/docs](https://github.com/r2day/docs/blob/main/architecture/2026-11-19-api-optimization.md)。

Mongo / Redis 连接和一批账号模型。多个 API 通过 `go.mod` 间接或直接使用 `github.com/r2day/db`。

## 问题

### P1 Redis 客户端不能传密码

`redis.go` 的 `InitRedisDB` 把 `Password` 写成空字符串。注释写成「Redis账号」，字段实际是密码。生产若按 `r2day/arch` 的 R03 给 Redis 加口令，这个函数连不上，调用方会继续留着无密码的 Redis。

### P1 密码字段会进 JSON

`models/AccountManage.go` 的 `Password` 是 `json:"password"`。`models/MerchantAccount.go` 同样。`GetOneByMerchantId` 旁的注释写了不要把内部字段返回给客户，结构体仍会序列化密码。本库没有哈希逻辑。调用方若把结构体直接写进响应，数据库里的密码会离开服务。

### P2 Mongo 连接失败直接 panic

`mongo.go` 的 `InitMongoDB` 在 URI 为空时 `log.Fatal`，`mongo.Connect` 失败时 `panic`。没有 Ping。

## 步骤

1. `InitRedisDB` 增加密码参数，空口令只允许本机开发。生产 Redis 加口令后，各服务一起改调用处。
2. `Password` 改为 `json:"-"`。响应使用不含密码的结构。确认库里存的是哈希而不是明文。
3. `InitMongoDB` 返回 `error`，连接后 Ping。不要在库里 `panic`。
