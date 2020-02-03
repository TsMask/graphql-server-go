package main

import (
	"encoding/json"
	"graphql-server-go/gplparse"
	"graphql-server-go/gplschema"
	"html/template"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	// 首页测试
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// HTML文本输出
		res.Header().Add("Content-Type", "text/html; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		// 模板文件
		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Printf("res:%v \n  err:%v \n", res, err)
			return
		}
		err = t.Execute(res, "GraphiQL")
	})

	// 请求入口
	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		// JSON格式输出，状态200
		res.Header().Add("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		// 解析请求参数
		opts := gplparse.ParseRequestOptions(req)
		// 进行graphql查询
		result := gplschema.ExecuteQuery(&graphql.Params{
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			Context:        req.Context(),
		})
		// 错误输出
		if len(result.Errors) > 0 {
			log.Printf("errors: %v", result.Errors)
		}
		// map转json序列化
		buff, _ := json.Marshal(result)
		_, _ = res.Write(buff)
	})

	log.Printf("Server is running on 127.0.0.1:8080 \n")
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
