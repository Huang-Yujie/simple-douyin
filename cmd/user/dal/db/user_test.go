package db

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
}

func TestCreateUser(t *testing.T) {
	Init()
	u := &User{}
	err := u.CreateUser(context.Background(), "123@163.com", "123456")
	if err != nil {
		t.Fatal(err.Error())
	}
	return
}

func TestGetUserById(t *testing.T) {
	Init()
	u := &User{}
	var id uint
	id = 1
	user, err := u.GetUserById(context.Background(), id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if user.ID != id {
		t.Fatal(errors.New("查询错误"))
	}
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	Init()
	id := 1
	email := "456@163.com"
	u := &User{}

	u1 := User{}
	u1.ID = uint(id)
	u1.Email = email

	err := u.UpdateUser(context.Background(), u1)
	if err != nil {
		t.Fatal(err.Error())
	}
}