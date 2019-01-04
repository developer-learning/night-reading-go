# Go 夜读

[![Build Status](https://travis-ci.org/developer-learning/night-reading-go.svg?branch=master)](https://travis-ci.org/developer-learning/night-reading-go) [![Go Report Card](https://goreportcard.com/badge/github.com/developer-learning/night-reading-go)](https://goreportcard.com/report/github.com/developer-learning/night-reading-go)  [![Documentation](https://godoc.org/github.com/developer-learning/night-reading-go?status.svg)](http://godoc.org/github.com/developer-learning/night-reading-go) [![Coverage Status](https://coveralls.io/repos/github/developer-learning/night-reading-go/badge.svg?branch=master)](https://coveralls.io/github/developer-learning/night-reading-go?branch=master) [![GitHub issues](https://img.shields.io/github/issues/developer-learning/night-reading-go.svg?label=Issue)](https://github.com/developer-learning/night-reading-go/issues) [![license](https://img.shields.io/github/license/developer-learning/night-reading-go.svg)](https://github.com/developer-learning/night-reading-go/blob/master/LICENSE) [![Release](https://img.shields.io/github/release/developer-learning/night-reading-go.svg?label=Release)](https://github.com/developer-learning/night-reading-go/releases) [![star this repo](http://githubbadges.com/star.svg?user=developer-learning&repo=night-reading-go)](http://github.com/developer-learning/night-reading-go) [![fork this repo](http://githubbadges.com/fork.svg?user=developer-learning&repo=night-reading-go)](http://github.com/developer-learning/night-reading-go/fork)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/developer-learning/night-reading-go.svg)](https://starcharts.herokuapp.com/developer-learning/night-reading-go)

[night-reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/?u=developer-learning&r=night-reading-go)

每周约定一个晚上进行 Go 源码阅读，Go 新手可以先去这里看看 **[Go 学习之路](https://github.com/developer-learning/learning-golang)**。

>阅读计划：Go 标准包、开源项目源代码。
>加油~
----

## [Go 夜读 YouTube 视频回放](https://www.youtube.com/channel/UCZwrjDu5Rf6O_CX2CVx7n8Q)

## 0. 下期预告

- **2018-06-21 net/http**

## 1. 往期回顾

| 源码总结 | YouTube |
|:----|:----|
| [2018-06-28-(net/http/server.go、request.go和net/textproto/reader.go)](./reading/20180628/README.md)| [https://www.youtube.com/watch?v=xodlVBWxTYM](https://www.youtube.com/watch?v=xodlVBWxTYM) |
| [2018-06-14-(net/http/server.go 和 h2_bundle.go 系列三)](./reading/20180614/README.md) | 无
| [2018-05-31-(net/http/server.go 系列二)](./reading/20180531/README.md) | [https://youtu.be/U84dn76gixQ](https://youtu.be/U84dn76gixQ)
| [2018-05-24-(net/http/server.go 系列一)](./reading/20180524/README.md) | [https://youtu.be/H3oXjpiOReQ](https://youtu.be/H3oXjpiOReQ)
| [2018-05-17-(strings/strings.go 系列二)](./reading/20180517/README.md)||
| [2018-05-10-(strings/strings.go 系列一)](./reading/20180510/README.md)||
| [2018-04-25-(strings/replace.go)、strings/search.go](./reading/20180425/README.md)||
| [2018-04-18-(strings/builder.go、strings/compare.go、strings/reader.go)](./reading/20180418/README.md) ||
| [2018-04-11-(telport、tp-micro、ants)](./reading/20180411/README.md) | |
| [2018-03-21-(cannot take address of temporary variables、telport、goutil、neochain)](./reading/20180321/README.md) | |

>[查看更多](./reading/README.md)

## 2. 日常讨论总结

- [2018-06-07 tcp连接设置timeout的问题讨论](./discuss/2018-06-07-dial-timeout-in-go.md)
- [2018-05-28 Go 语言中 CPU 和内存分析的问题讨论](./discuss/2018-05-28-pprof-in-go.md)
- [2018-05-23 锁失效和RPC框架选择的问题讨论](./discuss/2018-05-23-wechat-discuss.md)
- [2018-05-22 字符串转字节切片的问题讨论](./discuss/2018-05-22-go-string-to-byte-slice.md)
- [2018-05-21 在循环迭代值在 goroutine 中的使用等问题讨论](./discuss/2018-05-21-using-goroutines-on-loop-iterator-variables.md)
- [2018-05-18 bitset 实现和循环导入问题讨论](./discuss/2018-05-18-bitset-and-import-cycle-not-allowed.md)
- [2018-05-13 变量的作用域是贯穿整个 if-else-if 的](./discuss/2018-05-13-declaring-variables-on-if-else.md)
- [Go Vendor && GOPATH 讨论](./discuss/2018-05-10-which-vendor-tool.md)
- [2018-05-09 调试-开发工具-赋值:=和=的差别-变量作用域等等](./discuss/2018-05-09-wechat-discuss.md)
- [2018-05-08 Go 语言中下划线的分析总结](./discuss/2018-05-08-anlayze-underscore-in-go.md)

>[查看更多](./discuss/README.md)

## 3. 技术分析总结

- [2018-05-31 批量删除redis key](./articles/2018-05-31_batch-del-redis-key.md)

>[查看更多](./articles/README.md)

## 4. 深度剖析

1. 深度剖析 Boyer-Moore 和 Rabin-Karp 等字符串搜索算法。
2. 深度剖析 bitset 。

## 5. 大咖技术分享

有兴趣的可以联系我，并且提供你要分享的话题。

[其他更多](./other/README.md)

----

## 我们的目标

我们希望可以推进大家深入了解 Go ，快速成长为资深的 Gopher 。我们希望每次来了的人和没来的人都能够有收获，成长。

## 我们的方式

由一个主讲人带着大家一起去阅读 Go 源代码，一起去啃那些难啃的算法、学习代码里面的奇淫技巧，遇到问题或者有疑惑了，我们可以一起去检索，解答这些问题。我们可以一起学习，共同成长。

>阅读规则：选取 package 包，然后从上往下开始读 xxx.go 文件，每个文件从上往下读导出的函数（一步一步跟逻辑，如果逻辑跳出这个 package 则不做深入探究）。

## 我们的精神

开源！开源！开源！重要的事，一定要说三遍。

>希望有兴趣的小伙伴们一起加入，让我们一起把 《Go 夜读》建立成一个对大家都有帮助的开源社区。

## 怎么加入

目前微信群已经超过 100 人，请微信搜索 mai_yang ，然后备注你的姓名、公司、工作职责，来源：Github。

## 时间

20：00~22：00

## 地点

- 深圳市南山区桂庙路口向南瑞峰创业中心B座3079（南山地铁站E1出口，然后往前走 50 米）
