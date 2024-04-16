package graph

import (
	"context"
	"net/http"

	"example.com/federation-demo/users/graph/model"
	"github.com/vikstrous/dataloadgen"
)

const loadersKey = "dataloaders"

type userLoader struct {
	*dataloadgen.Loader[string, *model.User]
}

func NewUserLoader(repo *userRepo) *userLoader {
	return &userLoader{
		Loader: dataloadgen.NewLoader(repo.FindUsersByIDs),
	}
}

func UserLoaderFromContext(ctx context.Context) *userLoader {
	return ctx.Value(loadersKey).(*userLoader)
}

func UserLoaderMiddleware(repo *userRepo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewUserLoader(repo)
		ctx := context.WithValue(r.Context(), loadersKey, loader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
