package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
	"math/rand"
	"strconv"
)

func generateID() string {
	return strconv.Itoa(rand.Int())
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return []*model.Item{{
		ID:   generateID(),
		Name: "item",
		Category: &model.Category{
			ID:   generateID(),
			Name: "item_category",
		},
	}}, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return []*model.Category{{
		ID:   generateID(),
		Name: "category",
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
