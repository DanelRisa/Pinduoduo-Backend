package services

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUserByID(userID uint) (*User, error) {
	client := resty.New()

	resp, err := client.R().
		SetResult(&User{}).
		Get(fmt.Sprintf("http://localhost:8081/users/%d", userID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("user-service error: %s", resp.Status())
	}

	user := resp.Result().(*User)
	return user, nil
}
