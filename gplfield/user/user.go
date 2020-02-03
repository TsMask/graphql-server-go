package user

// 这个文件对应编写结构体
// 大写表示允许跨包调用，一些小写同包范围访问使用
// 这里不多做编写介绍了，模拟一段数据进行操作

// testUserList 固定测试用户数据
var testUserList = []map[string]interface{}{
	{
		"id":       "user-1",
		"username": "11Harry Potter and the Philosopher's Stone",
		"password": "11",
		"info":     "info-1",
	},
	{
		"id":       "user-2",
		"username": "22Harry Potter and the Philosopher's Stone",
		"password": "22",
		"info":     "info-2",
	},
	{
		"id":       "user-3",
		"username": "33Harry Potter and the Philosopher's Stone",
		"password": "33",
		"info":     "info-3",
	},
}

// testUserInfoList 固定测试用户信息数据
var testUserInfoList = []map[string]interface{}{
	{
		"id":        "info-1",
		"age":       "11",
		"firstName": "11Herman",
		"lastName":  "11Rowling",
	},
	{
		"id":        "info-2",
		"age":       "22",
		"firstName": "22Herman",
		"lastName":  "22Rowling",
	},
	{
		"id":        "info-3",
		"age":       "33",
		"firstName": "33Herman",
		"lastName":  "33Rowling",
	},
}
