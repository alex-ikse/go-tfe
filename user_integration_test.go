//go:build integration
// +build integration

package tfe

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersReadCurrent(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	u, err := client.Users.ReadCurrent(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, u.ID)
	assert.NotEmpty(t, u.AvatarURL)
	assert.NotEmpty(t, u.Username)

	t.Run("two factor options are decoded", func(t *testing.T) {
		assert.NotNil(t, u.TwoFactor)
	})
}

func TestUsersUpdate(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	uTest, err := client.Users.ReadCurrent(ctx)
	require.NoError(t, err)

	// Make sure we reset the current user when we're done.
	defer func() {
		_, err := client.Users.Update(ctx, UserUpdateOptions{
			Email:    String(uTest.Email),
			Username: String(uTest.Username),
		})
		if err != nil {
			t.Logf("Error updating user: %s", err)
		}
	}()

	t.Run("without any options", func(t *testing.T) {
		_, err := client.Users.Update(ctx, UserUpdateOptions{})
		require.NoError(t, err)

		u, err := client.Users.ReadCurrent(ctx)
		assert.NoError(t, err)
		assert.Equal(t, u, uTest)
	})

	t.Run("with a new username", func(t *testing.T) {
		_, err := client.Users.Update(ctx, UserUpdateOptions{
			Username: String("NewTestUsername"),
		})
		require.NoError(t, err)

		u, err := client.Users.ReadCurrent(ctx)
		assert.NoError(t, err)
		assert.Equal(t, "NewTestUsername", u.Username)
	})

	t.Run("with a new email address", func(t *testing.T) {
		_, err := client.Users.Update(ctx, UserUpdateOptions{
			Email: String("newtestemail@hashicorp.com"),
		})
		require.NoError(t, err)

		u, err := client.Users.ReadCurrent(ctx)

		email := ""
		if u.UnconfirmedEmail != "" {
			email = u.UnconfirmedEmail
		} else if u.Email != "" {
			email = u.Email
		} else {
			t.Fatalf("cannot test with user %q because both email and unconfirmed email are empty", u.ID)
		}

		assert.NoError(t, err)
		assert.Equal(t, "newtestemail@hashicorp.com", email)
	})

	t.Run("with invalid email address", func(t *testing.T) {
		u, err := client.Users.Update(ctx, UserUpdateOptions{
			Email: String("notamailaddress"),
		})
		assert.Nil(t, u)
		assert.Error(t, err)
	})
}

// This test is included in case you wish to run it during a review. Due to
// the trivial nature of this test, I'll remove it once the PR has been approved
func TestUsersFetchByID(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// include user id here
	u, err := client.Users.FetchByID(ctx, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, u.ID)
	assert.NotEmpty(t, u.AvatarURL)
	assert.NotEmpty(t, u.Username)

	t.Run("permissions are decoded", func(t *testing.T) {
		assert.NotNil(t, u.Permissions)
	})
}
