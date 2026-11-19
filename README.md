## 学习路线

### 1. Go 官方入门

- 从包、模块和 `go` 命令入手，了解如何编写、构建和测试 Go 程序。[How to Write Go Code](https://go.dev/doc/code)

### 2. 示例与中文视频

- 按顺序读一遍示例，在本地运行后改几个输入或条件，看看输出怎么变。学完本仓库的数组示例，可以接着练切片、Map 和函数。[Go by Example 中文版](https://gobyexample-cn.github.io/)
- 遇到指针、接口或并发等不熟悉的概念，可以配合 Gin-Vue-Admin 团队的中文视频学习。[Gin-Vue-Admin Go 教程](https://www.gin-vue-admin.com/guide/video/golang.html)

### 3. Web 后端实战

- 先用 Gin 写一个返回 JSON 的接口，再练习路由分组和请求参数绑定。[Gin 中文快速入门](https://gin-gonic.com/zh-cn/docs/quickstart/)
- 用 GORM 连接 SQLite 或 MySQL，把接口的数据写入数据库，练习增删改查。[GORM 连接数据库](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
- 写完小项目后，参考 Gin-Vue-Admin 的后台管理项目结构，了解前后端如何配合，以及认证和权限如何组织。[Gin-Vue-Admin 项目介绍](https://www.gin-vue-admin.com/guide/introduce/project.html)

### 4. 进阶阅读

- 根据学习思维导图选一个专题，结合示例阅读 Go 的底层实现。可以先选自己用过的 `context`、`channel` 或 `sync.Mutex`。[小徐先生的 gohub](https://github.com/xiaoxuxiansheng/gohub)

### 5. 英文课程

- Matt Holiday 的系列课程从基础类型讲到接口、并发、测试和性能分析。[Go Class 视频课程（YouTube）](https://www.youtube.com/watch?v=iDQAZEJK8lI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
- 课程讲义按章节排列，便于按主题复习。[Go Class 配套讲义（GitHub）](https://github.com/matt4biz/go-class-slides/blob/trunk/xmas-2020/README.md)

## 许可与来源

本仓库的原创内容采用 [MIT License](LICENSE)。

练习中的部分代码和注释改编自 Go by Example（Mark McGranaghan、Eli Bendersky 编写，gobyexample-cn 翻译），改动包括补充中文注释和调整部分示例输出。这些改编内容沿用上游的 CC BY 3.0 许可。[原版许可说明](https://github.com/mmcgrana/gobyexample#license) · [中文版](https://gobyexample-cn.github.io/) · [CC BY 3.0](https://creativecommons.org/licenses/by/3.0/)
