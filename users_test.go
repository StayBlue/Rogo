package Rogo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetUser(t *testing.T) {
	client := NewClient("")

	id := 1
	user, err := client.GetUser(id)
	assert.NoError(t, err, "err should be nil")

	assert.Equal(t, user.Id, id)
	assert.Equal(t, user.Name, "Roblox")
	assert.Equal(t, user.Verified, true)
}
