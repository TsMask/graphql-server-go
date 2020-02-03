# Graphql Go 基于Golang实践

>GraphQL 既是一种用于 API 的查询语言也是一个满足你数据查询的运行时。 GraphQL 对你的 API 中的数据提供了一套易于理解的完整描述，使得客户端能够准确地获得它需要的数据，而且没有任何冗余，也让 API 更容易地随着时间推移而演进，还能用于构建强大的开发者工具。

基于node的服务端开发中，GraphQL技术较为成熟常用。Golang作为高性能的现代语言在web服务器开发中也有很高的使用率，配合使用真香。

根据[GraphQL 官网代码]([https://graphql.cn/code/#go](https://graphql.cn/code/#go)
)中找到`graphql-go`：一个 Go/Golang 的 GraphQL 实现。

这个库还封装 [graphql-go-handler](https://github.com/graphql-go/graphql-go-handler)：通过HTTP请求处理GraphQL查询的中间件。

## 最后

通过`graphql-go`这个库实现GraphQl查询，使用`type`、`input`、`enum`的参数方式，实现CRUD操作。微服务采用`restful`和`graphql`两种方式进行开发，两者相辅相成，比如：上传、websocket等一些接口混用的模式。

> 示例环境 `go:1.13`，编辑器`vscode`    
> 项目下命令执行：`go run main.go`    
> GraphQL官网实现`JavaScript`的，使用`typescript`强类型代码编写。    
> 如果不会语法，建议先看看`JavaScript`基础SDL文件编写。    
