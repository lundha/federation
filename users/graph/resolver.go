package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repo *userRepo
}

func NewResolver() *Resolver {
	return &Resolver{Repo: &userRepo{limit: 25}}
}
