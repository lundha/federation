package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"example.com/federation-demo/suppliers/graph/model"
	faker "github.com/go-faker/faker/v4"
)

// GetSupplier is the resolver for the getSupplier field.
func (r *queryResolver) GetSupplier(ctx context.Context, id string) (*model.Supplier, error) {
	for _, supplier := range suppliers {
		if supplier.ID == id {
			return supplier, nil
		}
	}
	return nil, fmt.Errorf("supplier not found")
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var suppliers = initSuppliersData(10)

func initSuppliersData(count int) []*model.Supplier {
	countries := []string{"Norway", "Sweden", "UK", "Bahrain", "India"}
	suppliers := []*model.Supplier{}
	for i := 0; i < count; i++ {
		country := countries[i%len(countries)]
		suppliers = append(suppliers, &model.Supplier{
			ID:      fmt.Sprintf("supplier-%d", i+1),
			Name:    fmt.Sprintf("%s %s", faker.Word(), faker.Word()),
			Country: &country,
		})
	}
	return suppliers
}
