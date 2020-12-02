package service

import (
	"Go-000/Week02/biz"
	"fmt"
)

// 服务层
func ApiHasUserWithName(name string) (result bool, err error) {
	names, err := biz.GetAllUsersName()
	if err != nil {
		// 错误处理
		fmt.Printf("service: error: %+v\n", err)
		return
	}
	for _, n := range names {
		if n == name {
			result = true
			return
		}
	}
	return
}
