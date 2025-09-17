REST（Representational State Transfer）是一种基于 **Web** 的软件架构风格，旨在通过统一的接口实现资源的操作。

RESTful 风格以资源为中心，每个资源通过唯一的 **URI**（统一资源标识符）表示。资源的操作通过 HTTP 方法实现：

- **GET**：获取资源。
- **POST**：创建资源。
- **PUT**：更新资源。
- **DELETE**：删除资源。

RESTful 风格的设计遵循**无状态性**：每个请求都应包含完成操作所需的全部信息，服务器不保存客户端状态。

RESTful 一般将资源名作为 URL，而方法（GET/POST/PUT/DELETE）表示需要进行的操作。例如：

```
GET /usr/1001 #获取1001用户信息
POST /usr/1001 #创建1001用户信息

GET /usr/1001/create?pwd=1234 #错误示范，操作与资源名混杂
```

可以认为，URL 里**没有动词**，没有 `getUsr`、`addUsr`、`deleteUsr` 这类写法，否则就要检查是否违背规范。
