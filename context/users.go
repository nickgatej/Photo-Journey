package context

import (
	"context"
	"github.com/nickgatej/Photo-Journey/models"
)

type key string

const (
	userKey key = "user"
)

// Store User inside of a Context
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// Retrieve user from a Context
func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	user, ok := val.(*models.User)
	if !ok {
		// The most likely case is that nothing was ever stored in the context,
		// so it doesn't have a type of *models.User. It is also possible that
		// other code in this package wrote an invalid value using the user key.
		return nil
	}
	return user
}
