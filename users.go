package Rogo

import (
	"fmt"
	"github.com/carlmjohnson/requests"
	"time"
)

// MinimalUser represents a minimal user object.
// This is used for authentication purposes.
type MinimalUser struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type BaseUser struct {
	MinimalUser
	Verified       bool `json:"hasVerifiedBadge"`
	MembershipType int  `json:"buildersClubMembershipType"`
}

type User struct {
	BaseUser
	client      *Client
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Banned      bool      `json:"isBanned"`
}

func (c *Client) getUserRequest() *requests.Builder {
	return requests.
		New(c.config).
		BaseURL("https://users.roblox.com/v1/users/")
}

func (c *Client) GetUser(userId int) (*User, error) {
	var user User
	user.client = c

	if c.client == nil {
		return nil, fmt.Errorf("httpclient is nil")
	}

	err := c.getUserRequest().
		Path(fmt.Sprintf("%d", userId)).
		ToJSON(&user).
		Fetch(ctx)

	return &user, err
}

func (u *BaseUser) GetUser(c *Client) (*User, error) {
	return c.GetUser(u.Id)
}
