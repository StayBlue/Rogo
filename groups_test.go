package Rogo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var groupId = 10720185

func TestClient_GetGroup(t *testing.T) {
	client := NewClient("")

	id := groupId
	group, err := client.GetGroup(id)
	assert.NoError(t, err, "err should be nil")

	assert.Equal(t, group.Id, id, "group.Id should be equal to id")
	assert.Equal(t, group.Name, "Roblox Developer Community", "group.Name should be equal to 'Roblox Developer Community'")
	assert.Equal(t, group.Owner.Name, "InceptionTime", "group.Owner.Name should be equal to 'InceptionTime'")
}

func TestGroup_GetRoles(t *testing.T) {
	client := NewClient("")

	id := groupId
	group, err := client.GetGroup(id)
	assert.NoError(t, err, "err should be nil")

	roles, err := group.GetRoles()
	assert.NoError(t, err, "err should be nil")

	assert.Equal(t, len(roles), 5, "len(roles) should be equal to 5")
}

func TestGroup_GetRole(t *testing.T) {
	client := NewClient("")

	id := groupId
	group, err := client.GetGroup(id)
	assert.NoError(t, err, "err should be nil")

	roleNumber := 255
	role, err := group.GetRole(roleNumber)
	assert.NoError(t, err, "err should be nil")

	assert.Equal(t, role.Rank, roleNumber, "role.Rank should be equal to roleNumber")
}

func TestGroup_GetMembers(t *testing.T) {
	client := NewClient("")

	id := groupId
	group, err := client.GetGroup(id)
	assert.NoError(t, err, "err should be nil")

	members, err := group.GetMembers()
	assert.NoError(t, err, "err should be nil")

	assert.Equal(t, len(members.Data), 10)

	members, err = members.Next()
	assert.NoError(t, err, "err should be nil")

	user, err := members.Data[0].GetUser(client)
	assert.NoError(t, err, "err should be nil")

	assert.Greaterf(t, user.Id, 0, "user.Id should be greater than 0")
	assert.NotEqual(t, user.Name, "", "user.Name should not be empty")
}
