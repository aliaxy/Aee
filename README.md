# Go语言手写Web框架——Aee

## 介绍
Go 中 `net/http`提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。
一些Web开发中简单的需求并不支持，需要手工实现。
* 动态路由：例如hello/:name，hello/*这类的规则。
* 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
* 模板：没有统一简化的HTML机制。
* ……

用于学习目的，我实现了一个简单的Web框架，Aee，代码量只有1K，实现了基本的路由映射，中间件，模板渲染等功能。

Aee是使用 Go 语言实现一个简单的 Web 框架。
我第一次接触的 Go 语言的 Web 框架是Gin，Gin的代码总共是14K，其中测试代码9K，也就是说实际代码量只有5K。
Gin也是我非常喜欢的一个框架，与Python中的Flask很像，小而美。