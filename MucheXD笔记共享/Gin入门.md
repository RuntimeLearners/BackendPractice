
### 初始化并启动服务

使用Gin，首先需要初始化一个Gin引擎
```go
// gin.Default 函数会生成一个默认的 Engine （路由引擎）对象
// 集成了 Logger 和 Recovery 两个中间件
// Engine 是 Gin 框架最重要的数据结构，它是 Gin 框架的入口，本质上是一个 Http Handler
// 是一个用于处理 HTTP 请求的对象，维护一张路由表，将不同的 HTTP 请求路径映射到不同的处理函数上
r := gin.Default()
// 启动服务并监听,只接受一个 ip:port 格式的 string 参数，表示服务运行的 ip 地址和端口号
// ":8080" 是简写，省略了 ip，表示监听本地所有 ip 的 8080 端口，接收并处理 HTTP 请求
err := r.Run(":8080")
```
^bbe0df

### 添加路由表

在引擎中添加路由表以确定请求url最终处理的函数。

```go
//以下两种写法完全相同
r.GET("/hello",SayHello)
r.Handle("Get", "/hello", SayHello)
//对于不同的请求方法可以共用一个url
r.POST("/hello",SayHello)
//通配
r.Any("/hello",SayHello) //匹配所有方法的请求
r.NoRoute(WhatCanISay) //路由表中无对应项的处理
r.NoMethod(WharCanISay) //该资源路径有路由匹配，可是没有对应方法的路由匹配
```

上面代码中，`SayHello` 是一个 `gin.HandlerFunc`，在 Gin 里，能被强制转换成 gin.HandlerFunc 的只有下面这一种原型：

```go
func(*gin.Context)
```

也就是说，任何接受一个 gin.Context 参数且没有返回值的函数，都可以直接作为 `gin.HandlerFunc` 使用（或通过 `gin.HandlerFunc(f)` 显式转换）。
Gin **没有**为 HandlerFunc 定义其他重载或变体形式。

#### 路由组

路由组是方便开发者的功能，提供逻辑上的区分以及路径的简化。
- 可以抽离共同的 URL 前缀减短长度
- 可以在视觉上结构化路由表
- 编译后与无路由组的写法等效
- 路由组可以嵌套

```go
usrRoute := r.Group("/usr")//创建并获取路由组地址
//原型如下
func (group *gin.RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
//将路由添加到路由组
usrRoute.GET("/login",function)
```

### 中间件

Gin 允许开发者在请求路由到处理函数前先经过中间件（一种钩子函数）。中间件适合处理一些**公共的业务逻辑**，比如鉴权、日志等。
- 中间件是一个 `gin.HandlerFunc` 函数
- 中间件可以返回响应，因为传入了 `*gin.Context`
- 可以存在多个中间件
- 实际上中间件和处理函数在代码层面没有区别，只是逻辑上的区分
- 可以为路由组统一注册中间件，参考全局中间件写法

```go
//中间件可以直接添加到路由
r.GET("/usr/",midware,procfunc)
r.GET("/usr/",midware1,midware2,midware3,procfunc)//中间件和处理函数按序执行
//添加全局中间件，所有路由都会使用该中间件
//全局中间件间的执行顺序是添加顺序，全局中间件在局部中间件/处理函数前执行
r.Use(g_midware)
```

> 如果使用 [[Gin入门#^bbe0df|gin.Default()]] 生成引擎，则会默认存在 Logger 和 Recovery 两个中间件，如果不需要这两个中间件，请使用 `gin.New()`

执行中间件和处理函数的过程中，有一些联动各函数的功能。

控制中间件/处理函数的执行
```go
c.Abort() //停止后续中间件/处理函数，但当前中间件会执行直到完成
c.Next() //在当前位置先执行下一个中间件/处理函数，等待其完成后继续当前中间件/处理函数
```
> 得益于会返回到原函数的特性，`c.Next()` 常被用于计算处理函数的耗时

在中间件/处理函数间传递数据
```go
c.Set("key",val) //在上下文(gin.Context)中保存一个键值对
val, isExist := c.Get("key") //在上下文(gin.Context)中读取一个键对应的值
```
