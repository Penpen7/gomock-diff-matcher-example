package repository

import (
	"time"

	"github.com/Penpen7/gomock-diff-matcher-example/cmd"
)

type OnMemoryUserRepository struct {
	data map[int]cmd.User
}

func NewOnMemoryUserRepository() *OnMemoryUserRepository {
	data := map[int]cmd.User{
		1: {ID: "1", Name: "Alice", CreatedAt: time.Now()},
		2: {ID: "2", Name: "Bob", CreatedAt: time.Now()},
	}
	return &OnMemoryUserRepository{data}
}

func (r *OnMemoryUserRepository) Create(user cmd.User) error {
	r.data[len(r.data)+1] = user
	return nil
}

func (r *OnMemoryUserRepository) FindByID(id int) *cmd.User {
	user, ok := r.data[id]
	if !ok {
		return nil
	}
	return &user
}
