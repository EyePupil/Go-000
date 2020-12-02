package biz

import "Go-000/Week02/dao"

func GetAllUsersName() (names []string, err error) {
	users, err := dao.NewUserDao().GetAllUsers()
	if err != nil {
		return
	}
	// do business
	for range users {
	}
	return
}
