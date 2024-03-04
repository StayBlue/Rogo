package Rogo

import (
	"fmt"
	"time"
)

type GroupUser struct {
	BaseUser
	Id   int    `json:"userId"`
	Name string `json:"username"`
}

type BaseGroup struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       GroupUser `json:"owner"`
	Shout       struct {
		Body    string    `json:"body"`
		Poster  GroupUser `json:"poster"`
		Created time.Time `json:"created"`
		Updated time.Time `json:"updated"`
	}
	Count    int  `json:"memberCount"`
	Locked   bool `json:"isLocked"`
	Verified bool `json:"hasVerifiedBadge"`
}

type Group struct {
	BaseGroup
	client *Client
}

type BaseRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rank        int    `json:"rank"`
	Count       int    `json:"memberCount"`
}

type Roles struct {
	GroupId int        `json:"groupId"`
	Roles   []BaseRole `json:"roles"`
}

type BaseMember struct {
	User GroupUser `json:"user"`
	Role BaseRole  `json:"role"`
}

func (c *Client) GetGroup(groupId int) (*Group, error) {
	var group Group
	group.client = c

	if c.client == nil {
		return nil, fmt.Errorf("httpclient is nil")
	}

	err := c.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d", groupId)).
		ToJSON(&group).
		Fetch(ctx)

	return &group, err
}

func (g *Group) getClient() (*Client, error) {
	client := g.client
	if client == nil {
		return nil, fmt.Errorf("group client is nil")
	}
	return client, nil
}

func (g *Group) GetRoles() ([]BaseRole, error) {
	client, err := g.getClient()
	if err != nil {
		return nil, err
	}

	var roles Roles
	err = client.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d/roles", g.Id)).
		ToJSON(&roles).
		Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return roles.Roles, err
}

func (g *Group) GetRole(roleNumber int) (*BaseRole, error) {
	roles, err := g.GetRoles()
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		var result bool
		if roleNumber <= 255 {
			result = role.Rank == roleNumber
		} else {
			result = role.Id == roleNumber
		}

		if result {
			return &role, nil
		}
	}

	return nil, fmt.Errorf("role not found")
}

func (g *Group) SetRank(userId int, roleId int) error {
	client, err := g.getClient()
	if err != nil {
		return err
	}

	err = client.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d/users/%d", g.Id, userId)).
		BodyJSON(map[string]int{
			"roleId": roleId,
		}).
		Patch().
		Fetch(ctx)

	return err
}

func (g *Group) Exile(userId int) error {
	client, err := g.getClient()
	if err != nil {
		return err
	}

	err = client.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d/users/%d", g.Id, userId)).
		Delete().
		Fetch(ctx)

	return err
}

func (g *Group) Join() error {
	client, err := g.getClient()
	if err != nil {
		return err
	}

	err = client.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d/users", g.Id)).
		Post().
		Fetch(ctx)

	return err
}

func (g *Group) GetMembers() (*Pagination[BaseMember], error) {
	client, err := g.getClient()
	if err != nil {
		return nil, err
	}

	var members Pagination[BaseMember]
	members.client = client
	members.URL = fmt.Sprintf("https://groups.roblox.com/v1/groups/%d/users", g.Id)
	err = client.getRequest("groups").
		Path(fmt.Sprintf("v1/groups/%d/users", g.Id)).
		ToJSON(&members).
		Fetch(ctx)

	return &members, err
}

func (m *BaseMember) GetUser(c *Client) (*User, error) {
	return c.GetUser(m.User.Id)
}
