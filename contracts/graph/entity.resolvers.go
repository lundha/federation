package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"example.com/federation-demo/contracts/graph/model"
)

// FindContractByID is the resolver for the findContractByID field.
func (r *entityResolver) FindContractByID(ctx context.Context, id string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: FindContractByID - findContractByID"))
}

// FindSupplierByID is the resolver for the findSupplierByID field.
func (r *entityResolver) FindSupplierByID(ctx context.Context, id string) (*model.Supplier, error) {
	fmt.Printf("    contracts-resolver:    contracts for supplier %s\n", id)
	var result []*model.Contract
	for _, c := range contracts {
		c := c
		if c.Supplier.ID == id {
			result = append(result, c)
		}
	}
	// limit for pretty output during demo
	result = result[:3]
	return &model.Supplier{ID: id, Contracts: result}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }