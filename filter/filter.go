package filter

// Request 输入 (空接口别名-任意类型)
type Request interface{}

// Response 输出 (空接口别名-任意类型)
type Response interface{}

// Filter 处理默认实现
type Filter interface {
	Process(data Request) (Response, error)
}

/*

处理流程：
	    "1，2，3"
split_filter  ["1","2","3"]
to_int_filter [1,2,3]
sum_filter    9
*/

/*
1. 验证手机号是注册登录
2. 验证手机号是完善用户信息: doctor表是否存在、
*/
