//go:generate mockgen -source=cmd.go -destination=../mock_cmd/cmd.go
package cmd

import (
	"fmt"
	"time"
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type UserRepository interface {
	FindByID(id int) *User
	Create(user User) error
}

func Run(userRepository UserRepository, id int) error {
	user := userRepository.FindByID(id)
	if user == nil {
		return fmt.Errorf("user not found")
	}

	user.Name = "Penpen7"

	if err := userRepository.Create(*user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
