package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"context"
	"log"
	"math/rand"
	"strconv"
)

func (r *queryResolver) Items(ctx context.Context) ([]*generated.Item, error) {
	log.Println("queryResolver: Items")
	return []*generated.Item{{
		ID:   generateID(),
		Name: "item",
		Category: &generated.Category{
			ID:   generateID(),
			Name: "item_category",
		},
	}}, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*generated.Category, error) {
	log.Println("queryResolver: Categories")
	return []*generated.Category{{
		ID:   generateID(),
		Name: "category",
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func generateID() string {
	return strconv.Itoa(rand.Int())
}
