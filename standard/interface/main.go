package main

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type IUser interface {
	GetUser(id int) (*User, error)
}

type UserService struct {
	users map[int]*User
}

func NewUser() IUser {
	// Q. ここアドレス返すのはなんで?
	// ポインタレシーバーでGetUserを実装しているから
	return &UserService{
		users: map[int]*User{
			1: {ID: 1, Name: "John"},
			2: {ID: 2, Name: "Jane"},
		},
	}
}

// 1
// ポインタレシーバの実装をすることで、
// UserServiceのインスタンスを直接操作できるようにする
// アドレスで参照しているUserオブジェクトに対してしか使えない
func (s *UserService) GetUser(id int) (*User, error) {
	if user, ok := s.users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}