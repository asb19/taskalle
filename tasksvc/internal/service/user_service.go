package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asb19/tasksvc/internal/model"
	"github.com/google/uuid"
)

type UserClient struct {
	BaseURL string
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{BaseURL: baseURL}
}

func (c *UserClient) GetUser(id uuid.UUID) (model.User, error) {
	url := fmt.Sprintf("%s/users/%s", c.BaseURL, id.String())
	resp, err := http.Get(url)
	if err != nil {
		return model.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.User{}, fmt.Errorf("failed to fetch user: %s", resp.Status)
	}

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return model.User{}, err
	}
	fmt.Printf(user.Name)

	return user, nil
}
