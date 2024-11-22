package handler

import (
	"fmt"
	"io"
	"net/http"
)

type UserServiceGateway struct {
	BaseURL string
}

func NewUserServiceGateway(baseURL string) *UserServiceGateway {
	return &UserServiceGateway{BaseURL: baseURL}
}

func (c *UserServiceGateway) GetUser(userID string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s", c.BaseURL, userID))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
