package dao

import (
	"Go-000/Week02/model"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// DAO层
type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}


func (dao *UserDao) GetAllUsers() (users []model.User, err error) {
	data, err := dbGetUsers()
	if err == sql.ErrNoRows {
		// 未查询到在DAO进行处理后不往上抛
		fmt.Println("There are no eligible users")
		err = nil
		return
	}
	if err != nil {
		err = errors.Wrap(err, "failed query users err")
		return
	}
	// 填充数据到users切片中
	for range data {
		continue
	}
	return
}

func dbGetUsers() ([]map[string]interface{}, error) {
	return nil, sql.ErrNoRows
}
