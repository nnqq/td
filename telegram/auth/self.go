package auth

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/nnqq/td/tg"
)

// self returns current user.
//
// You can use tg.User.Bot to check whether current user is bot.
func (c *Client) self(ctx context.Context) (*tg.User, error) {
	users, err := c.api.UsersGetUsers(ctx, []tg.InputUserClass{&tg.InputUserSelf{}})
	if err != nil {
		return nil, err
	}

	user, ok := tg.UserClassArray(users).FirstAsNotEmpty()
	if !ok {
		return nil, xerrors.Errorf("users response count: %v", users)
	}

	return user, nil
}
