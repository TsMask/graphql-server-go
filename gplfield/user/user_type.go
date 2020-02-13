package user

import "github.com/graphql-go/graphql"

// queryType 用户类型
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "用户字段",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "ID值",
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "用户名",
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: "密码",
		},
		"info": &graphql.Field{
			Type:        infoType,
			Description: "关联用户信息",
			Resolve:     info, // 级联函数调用
		},
	},
})

// infoType 用户信息类型
var infoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserInfo",
	Description: "用户信息字段",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "ID值",
		},
		"age": &graphql.Field{
			Type:        graphql.Int,
			Description: "年龄",
		},
		"firstName": &graphql.Field{
			Type:        graphql.String,
			Description: "名",
		},
		"lastName": &graphql.Field{
			Type:        graphql.String,
			Description: "姓",
		},
	},
})

// lastNameEnum 姓-枚举类型
var lastNameEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "LastNameEnum",
	Description: "名选择",
	Values: graphql.EnumValueConfigMap{
		"OK": &graphql.EnumValueConfig{
			Value:       graphql.String,
			Description: "好的",
		},
		"SOLO": &graphql.EnumValueConfig{
			Value:       graphql.String,
			Description: "solo",
		},
	},
})

// createInput  用户创建修改-Mutation提交输入类型
var createInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "用户名",
		},
		"password": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "密码",
		},
		"info": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(infoCreateInput),
			Description: "关联用户信息",
		},
	},
})

// infoCreateInput 部分用户信息创建修改-Mutation提交输入类型
var infoCreateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInfoCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"age": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "年龄",
		},
		"firstName": &graphql.InputObjectFieldConfig{
			Type:        lastNameEnum,
			Description: "名",
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "姓",
		},
	},
})
