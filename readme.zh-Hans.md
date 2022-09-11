# kook

kook 是一个低级的和开黑啦机器人交互的软件库。目前她还在开发中。

[![Go Reference](https://pkg.go.dev/badge/github.com/lonelyevil/kook.svg)](https://pkg.go.dev/github.com/lonelyevil/kook)
[![Go Report Card](https://goreportcard.com/badge/github.com/lonelyevil/kook)](https://goreportcard.com/report/github.com/lonelyevil/kook)
[![Server Badger](https://img.shields.io/badge/%E5%BC%80%E9%BB%91%E5%95%A6-%E4%BA%A4%E6%B5%81%E7%BE%A4-informational)](https://kaihei.co/r5s1WO)


## 开始

直到 `v1.0.0` 发布为止，不建议在生产中使用。

### 安装

假定您有可用的 Go 开发环境。

```go get github.com/lonelyevil/kook```

除了库本身，你还需要安装一个结构化日志库和相应的适配器。
目前，只有 [phuslu/log](https://github.com/phuslu/log) 的适配器被实现了。
所以，想要使用本库，以下的安装命令也是必须的。

```go get github.com/lonelyevil/kook/log_adapter/plog```

### 使用

参考`examples`目录下的例子。

## 功能和路线

请阅读 [features.md](features.md)。

## 文档

正在编写中。

针对那些没有良好文档的代码，用户可以参考 [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) 。

## 版本号

在 `v1.0.0` 之前，我不能保证 API 的稳定。
在此之后，会遵循 [语义化版本号](https://semver.org/) 。


## 贡献

目前，这个仓库暂时不接受较大的 PR ，API 可能会在 PR 审核的过程中改变。
非常欢迎臭虫报告、修复以及代码上的建议。

## 授权协议

kook 是以 ISC 授权条款分发的一个免费开源的软件库。

## 特别感谢

<img alt="GoLand logo" src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.png" width="250">