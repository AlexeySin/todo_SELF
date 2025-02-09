package service

import (
	"context"
	entities "todo_SELF/auth/pkg/entities"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Register(ctx context.Context, creds entities.Credentials) (err error) {
	defer func() {
		l.logger.Log("method", "Register", "creds", creds, "err", err)
	}()
	return l.next.Register(ctx, creds)
}
func (l loggingMiddleware) Login(ctx context.Context, creds entities.Credentials) (key entities.Key, err error) {
	defer func() {
		l.logger.Log("method", "Login", "creds", creds, "key", key, "err", err)
	}()
	return l.next.Login(ctx, creds)
}
func (l loggingMiddleware) Access(ctx context.Context, key entities.Key) (e0 entities.Key, e1 error) {
	defer func() {
		l.logger.Log("method", "Access", "key", key, "e0", e0, "e1", e1)
	}()
	return l.next.Access(ctx, key)
}
func (l loggingMiddleware) Logout(ctx context.Context, key entities.Key) (e1 error) {
	defer func() {
		l.logger.Log("method", "Logout", "key", key, "e1", e1)
	}()
	return l.next.Logout(ctx, key)
}
func (l loggingMiddleware) UserRegistrationAttempt(ctx context.Context, creds entities.Credentials) (err error) {
	defer func() {
		l.logger.Log("method", "UserRegistrationAttempt", "creds", creds, "err", err)
	}()
	return l.next.UserRegistrationAttempt(ctx, creds)
}

func (l loggingMiddleware) FetchUsers(ctx context.Context, key entities.Key) (users []entities.User, err error) {
	defer func() {
		l.logger.Log("method", "FetchUsers", "key", key, "users", users, "err", err)
	}()
	return l.next.FetchUsers(ctx, key)
}

func (l loggingMiddleware) BlockUser(ctx context.Context, key entities.Key) (err error) {
	defer func() {
		l.logger.Log("method", "BlockUser", "key", key, "err", err)
	}()
	return l.next.BlockUser(ctx, key)
}
func (l loggingMiddleware) UnblockUser(ctx context.Context, key entities.Key) (err error) {
	defer func() {
		l.logger.Log("method", "UnblockUser", "key", key, "err", err)
	}()
	return l.next.UnblockUser(ctx, key)
}
func (l loggingMiddleware) SearchUsers(ctx context.Context, key entities.Key) (users []entities.User, err error) {
	defer func() {
		l.logger.Log("method", "SearchUsers", "key", key, "err", err)
	}()
	return l.next.SearchUsers(ctx, key)
}
func (l loggingMiddleware) DropUser(ctx context.Context, key entities.Key) (err error) {
	defer func() {
		l.logger.Log("method", "DropUser", "key", key, "err", err)
	}()
	return l.next.DropUser(ctx, key)
}
func (l loggingMiddleware) UpdateUser(ctx context.Context, user entities.User, key entities.Key) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateUser", "key", key, "user", user, "err", err)
	}()
	return l.next.UpdateUser(ctx, user, key)
}
