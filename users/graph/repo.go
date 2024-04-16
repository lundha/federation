package graph

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"example.com/federation-demo/users/graph/model"
)

type userRepo struct {
	concurrency int32
	limit       int32
}

func (r *userRepo) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	r.databaseOperation(ctx, id)
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")

}
func (r *userRepo) FindUsersByIDs(ctx context.Context, ids []string) ([]*model.User, []error) {
	r.databaseOperation(ctx, ids...)
	var result []*model.User
	var errors []error
	for _, id := range ids {
		id := id
		found := false
		for _, user := range users {
			if user.ID == id {
				result = append(result, user)
				found = true
			}
		}
		if !found {
			errors = append(errors, fmt.Errorf("user not found: %s", id))
		}
	}
	return result, errors
}

func (r *userRepo) databaseOperation(ctx context.Context, ids ...string) {
	atomic.AddInt32(&r.concurrency, 1)
	defer atomic.AddInt32(&r.concurrency, -1)
	time.Sleep(time.Millisecond * time.Duration(100+(time.Now().UnixNano()%400)))
	var s string
	if len(ids) == 1 {
		s = ids[0]
	} else {
		s = fmt.Sprintf("for %d users", len(ids))
	}
	fmt.Printf("    users-resolver:    database access [%s]\n", s)
	if r.concurrency >= r.limit {
		panic("PSQLException: FATAL: too many open connections")
	}
}
